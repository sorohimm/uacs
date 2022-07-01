package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/models"
	"uacs/internal/services"
)

type Controllers struct {
	Log       *zap.SugaredLogger
	Services  *services.Services
	Validator *validator.Validate
}

func (c *Controllers) NewCompetition(ctx *gin.Context) {
	var newCompetition models.Competition

	err := ctx.BindJSON(&newCompetition)
	if err != nil {
		c.Log.Error("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	competition, err := c.Services.NewCompetition(newCompetition)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, competition)
}
