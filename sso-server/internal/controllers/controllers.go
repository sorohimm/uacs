package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/models"
	"uacs/sso-server/internal/services"
)

type Controllers struct {
	Log      *zap.SugaredLogger
	Cfg      *config.Config
	Services *services.Services
}

func (c *Controllers) Registration(ctx *gin.Context) {
	var newUser models.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.Services.Registration(newUser)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *Controllers) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest

	err := ctx.BindJSON(&loginReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	loginResp, err := c.Services.Login(loginReq)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie("accessToken", loginResp.AccessToken, loginResp.ExpiresIn, "", "", true, true)
	ctx.SetCookie("refreshToken", loginResp.RefreshToken, loginResp.ExpiresIn, "", "", true, true)
	ctx.Status(http.StatusOK)
}

func (c *Controllers) Logout(ctx *gin.Context) {
	var session models.Session

	cookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	session.RefreshToken = cookie

	err = c.Services.Logout(session)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controllers) ValidateAccessToken(ctx *gin.Context) {
	var session models.Session

	cookie, err := ctx.Cookie("accessToken")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	session.AccessToken = cookie

	ok, err := c.Services.ValidateAccessToken(session)
	if err != nil || !ok {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controllers) GetUserId(ctx *gin.Context) {
	var session models.Session

	cookie, err := ctx.Cookie("accessToken")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	session.AccessToken = cookie

	userId, err := c.Services.GetUserId(session)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": userId})
}
