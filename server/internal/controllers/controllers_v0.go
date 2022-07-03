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

type ControllersV0 struct {
	Log        *zap.SugaredLogger
	ServicesV0 *services.ServicesV0
	Validator  *validator.Validate
}

func (c *ControllersV0) NewCompetition(ctx *gin.Context) {
	var newCompetition models.Competition

	err := ctx.BindJSON(&newCompetition)
	if err != nil {
		c.Log.Error("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	competition, err := c.ServicesV0.NewCompetition(newCompetition)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, competition)
}

func (c *ControllersV0) GetMyCompetitionsShort(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

func (c *ControllersV0) GetAllCompetitionsShort(ctx *gin.Context) {
	competitions, err := c.ServicesV0.GetAllCompetitionsShort()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, competitions)
}

func (c *ControllersV0) GetSingleCompetitionFull(ctx *gin.Context) {
	id := ctx.Query("id")
	competition, err := c.ServicesV0.GetSingleCompetitionFull(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, competition)
}
