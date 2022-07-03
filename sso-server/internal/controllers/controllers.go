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

	ctx.JSON(http.StatusCreated, "")
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
	ctx.JSON(http.StatusOK, loginResp)
}
