package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
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

	err := json.NewDecoder(ctx.Request.Body).Decode(&newCompetition)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	competition, err := c.ServicesV0.NewCompetition(newCompetition)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
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
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, competitions)
}

func (c *ControllersV0) GetSingleCompetitionFull(ctx *gin.Context) {
	id := ctx.Param("id")
	competition, err := c.ServicesV0.GetSingleCompetitionFull(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, competition)
}

func (c *ControllersV0) AddParticipant(ctx *gin.Context) {
	var newParticipant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&newParticipant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	participant, err := c.ServicesV0.AddParticipant(newParticipant)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusCreated, participant)
}

func (c *ControllersV0) AddJudge(ctx *gin.Context) {
	var newJudge models.CompetitionJudge

	err := json.NewDecoder(ctx.Request.Body).Decode(&newJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.ServicesV0.AddJudge(newJudge)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusCreated, judge)
}

func (c *ControllersV0) DeleteParticipant(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.ServicesV0.DeleteParticipant(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ControllersV0) DeleteJudge(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.ServicesV0.DeleteJudge(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ControllersV0) UpdateParticipant(ctx *gin.Context) {
	var updateParticipant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateParticipant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	participant, err := c.ServicesV0.UpdateParticipant(updateParticipant)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, participant)
}

func (c *ControllersV0) UpdateJudge(ctx *gin.Context) {
	var updateJudge models.CompetitionJudge

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.ServicesV0.UpdateJudge(updateJudge)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, judge)
}

func (c *ControllersV0) DeleteCompetition(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("empty id"))
		return
	}

	err := c.ServicesV0.DeleteCompetition(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

}

func (c *ControllersV0) UpdateCompetition(ctx *gin.Context) {
	var updateCompetition models.Competition

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateCompetition)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	competition, err := c.ServicesV0.UpdateCompetition(updateCompetition)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, competition)
}
