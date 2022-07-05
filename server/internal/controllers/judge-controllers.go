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
	var newJudge models.CompetitionJudge

	err := json.NewDecoder(ctx.Request.Body).Decode(&newJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.JudgeServices.AddJudge(newJudge)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusCreated, judge)
}

func (c *JudgeControllers) UpdateJudge(ctx *gin.Context) {
	var updateJudge models.CompetitionJudge

	err := json.NewDecoder(ctx.Request.Body).Decode(&updateJudge)
	if err != nil {
		c.Log.Errorf("Error occurred during unmarshalling. Error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	judge, err := c.JudgeServices.UpdateJudge(updateJudge)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, judge)
}

func (c *JudgeControllers) GetJudges(ctx *gin.Context) {

}

func (c *JudgeControllers) DeleteJudge(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.JudgeServices.DeleteJudge(id)
	if err != nil {
		ctx.AbortWithError(errStatusCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
