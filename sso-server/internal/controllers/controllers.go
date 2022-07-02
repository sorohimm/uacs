package controllers

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

	token, err := c.KeyloackClient.LoginAdmin(context.Background(), c.Cfg.KeycloakAdminUsername, c.Cfg.KeycloakAdminPassword, c.Cfg.KeycloakRealmName)
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

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = c.KeyloackClient.SetPassword(context.Background(), token.AccessToken, usrId, c.Cfg.KeycloakRealmName, string(hashPwd), false)
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

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(loginReq.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt, err := c.KeyloackClient.Login(context.Background(),
		c.Cfg.KeycloakClientId,
		c.Cfg.KeycloakClientSecret,
		c.Cfg.KeycloakRealmName,
		loginReq.Username,
		string(hashPwd),
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
