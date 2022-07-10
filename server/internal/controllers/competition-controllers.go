package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/models"
	"uacs/internal/services"
)

type CompetitionControllers struct {
	CompetitionServices *services.CompetitionServicesV0
	Log                 *zap.SugaredLogger
}

func (c *CompetitionControllers) NewCompetition(ctx *gin.Context) {
	var newCompetition models.Competition

	err := json.NewDecoder(ctx.Request.Body).Decode(&newCompetition)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	competition, err := c.CompetitionServices.CreateCompetition(newCompetition)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusCreated, competition)
}

func (c *CompetitionControllers) GetMyCompetitionsShort(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

func (c *CompetitionControllers) GetAllCompetitionsShort(ctx *gin.Context) {
	competitions, err := c.CompetitionServices.GetAllCompetitionsShort()
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, competitions)
}

func (c *CompetitionControllers) GetSingleCompetitionFull(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	competition, err := c.CompetitionServices.GetSingleCompetitionFull(id)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, competition)
}

func (c *CompetitionControllers) DeleteCompetition(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := c.CompetitionServices.DeleteCompetition(id)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}
}

func (c *CompetitionControllers) UpdateCompetition(ctx *gin.Context) {
	var updateCompetition models.Competition

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateCompetition)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	competition, err := c.CompetitionServices.UpdateCompetition(updateCompetition)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, competition)
}
