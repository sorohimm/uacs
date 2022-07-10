package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/models"
	"uacs/internal/services"
)

type JudgeControllers struct {
	JudgeServices *services.JudgeServicesV0
	Log           *zap.SugaredLogger
}

func (c *JudgeControllers) AddJudge(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	var newJudge models.CompetitionJudge
	err := json.NewDecoder(ctx.Request.Body).Decode(&newJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.JudgeServices.CreateJudge(competitionId, newJudge)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusCreated, judge)
}

func (c *JudgeControllers) UpdateJudge(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	var updateJudge models.CompetitionJudge
	err := json.NewDecoder(ctx.Request.Body).Decode(&updateJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.JudgeServices.UpdateJudge(competitionId, updateJudge)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.JSON(http.StatusOK, judge)
}

func (c *JudgeControllers) GetJudges(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	if competitionId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	judgeId := ctx.Param("id")

	var err error
	var result interface{}
	switch judgeId {
	case "":
		result, err = c.JudgeServices.GetJudges(competitionId)
	default:
		result, err = c.JudgeServices.GetJudge(competitionId, judgeId)
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *JudgeControllers) DeleteJudge(ctx *gin.Context) {
	competitionId := ctx.Query("competition")
	id := ctx.Param("id")
	if competitionId == "" || id == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	err := c.JudgeServices.DeleteJudge(competitionId, id)
	if err != nil {
		ctx.AbortWithStatus(errStatusCode(err))
		return
	}

	ctx.Status(http.StatusOK)
}
