package controllers

import (
	"context"
	"crypto/sha256"
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/models"
)

type Controllers struct {
	Log            *zap.SugaredLogger
	KeyloackClient gocloak.GoCloak
	Cfg            *config.Config
}

func (c *Controllers) Registration(ctx *gin.Context) {
	var newUser models.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := c.KeyloackClient.LoginAdmin(context.Background(), c.Cfg.KeycloakAdminUsername, c.Cfg.KeycloakAdminPassword, "master")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user := gocloak.User{
		FirstName: gocloak.StringP(newUser.FirstName),
		LastName:  gocloak.StringP(newUser.LastName),
		Email:     gocloak.StringP(newUser.Email),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(newUser.Username),
	}

	usrId, err := c.KeyloackClient.CreateUser(context.Background(), token.AccessToken, c.Cfg.KeycloakRealmName, user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	bytePwd := sha256.Sum256([]byte(newUser.Password))

	err = c.KeyloackClient.SetPassword(context.Background(), token.AccessToken, usrId, c.Cfg.KeycloakRealmName, string(bytePwd[:]), false)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}

func (c *Controllers) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest

	err := ctx.BindJSON(&loginReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bytePwd := sha256.Sum256([]byte(loginReq.Password))

	jwt, err := c.KeyloackClient.Login(context.Background(),
		c.Cfg.KeycloakClientId,
		c.Cfg.KeycloakClientSecret,
		c.Cfg.KeycloakRealmName,
		loginReq.Username,
		string(bytePwd[:]),
	)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	resp := models.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	ctx.JSON(http.StatusOK, resp)
}
