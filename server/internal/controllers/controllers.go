package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"uacs/internal/services"
)

type Controllers struct {
	Log      *zap.SugaredLogger
	Services *services.Services
}

func (c *Controllers) CreateUser(ctx *gin.Context) {

}
