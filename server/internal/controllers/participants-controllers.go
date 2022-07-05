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
	var newParticipant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&newParticipant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	participant, err := c.ParticipantServices.AddParticipant(newParticipant)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusCreated, participant)
}

func (c *ParticipantControllers) DeleteParticipant(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.ParticipantServices.DeleteParticipant(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ParticipantControllers) UpdateParticipant(ctx *gin.Context) {
	var updateParticipant models.CompetitionParticipant

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateParticipant)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	participant, err := c.ParticipantServices.UpdateParticipant(updateParticipant)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, participant)
}

func (c *ParticipantControllers) GetParticipants(ctx *gin.Context) {

}
