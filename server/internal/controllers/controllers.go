package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"uacs/internal/services"
)

type Controllers struct {
	Log       *zap.SugaredLogger
	Services  *services.Services
	Validator *validator.Validate
}

func (c *Controllers) NewCompetition(ctx *gin.Context) {

}
