package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/models"
	"uacs/internal/services"
)

type ParticipantControllers struct {
	ParticipantServices *services.ParticipantServicesV0
	Log                 *zap.SugaredLogger
}

func (c *ParticipantControllers) AddParticipant(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	var newParticipant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&newParticipant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	participant, err := c.ParticipantServices.CreateParticipant(competitionId, newParticipant)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusCreated, participant)
}

func (c *ParticipantControllers) DeleteParticipant(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	id := ctx.Param("id")
	if competitionId == "" || id == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	err := c.ParticipantServices.DeleteParticipant(competitionId, id)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ParticipantControllers) UpdateParticipant(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	var participant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&participant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	result, err := c.ParticipantServices.UpdateParticipant(competitionId, participant)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *ParticipantControllers) GetParticipants(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	participantId := ctx.Param("id")

	var err error
	var result interface{}
	switch participantId {
	case "":
		result, err = c.ParticipantServices.GetParticipants(competitionId)
	default:
		result, err = c.ParticipantServices.GetParticipant(competitionId, participantId)
	}
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
