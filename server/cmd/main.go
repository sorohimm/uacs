package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"time"
	"uacs/internal/config"
	"uacs/internal/infrastructure"
)

var (
	log *zap.SugaredLogger
	cfg *config.Config
	ctx context.Context
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("Loading logger error: %s\n", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()

	cfg, err = config.New()
	if err != nil {
		log.Fatalf("Config init error: %s\n", err)
	}
	log.Infof("Config loaded:\n%+v", cfg)

	ctx = context.Background()
}

func main() {
	gin.SetMode(gin.DebugMode) // TODO: switch on ReleaseMode on production

	injector, err := infrastructure.Injector(log, ctx, cfg)
	if err != nil {
		log.Fatalf("Ijection fatal error: %s\n", err.Error())
	}

	competitionsControllers := injector.InjectCompetitionsControllers()
	participantControllers := injector.InjectParticipantControllers()
	judgeControllers := injector.InjectJudgeControllers()
	middlewareV0 := injector.InjectMiddlewareV0()

	r := gin.Default()

	// Public handles
	r.GET("/competitions", competitionsControllers.GetAllCompetitionsShort)
	r.GET("/competitions/:id", competitionsControllers.GetSingleCompetitionFull)
	r.GET("/participants", participantControllers.GetParticipants)
	r.GET("/participants/:id", participantControllers.GetParticipants)
	r.GET("/judges", judgeControllers.GetJudges)
	r.GET("/judges/:id", judgeControllers.GetJudges)

	authorized := r.Group("/")
	authorized.Use(middlewareV0.AuthRequired)
	{
		edit := authorized.Group("/edit")
		{
			stuff := edit.Group("/stuff")
			stuff.Use() // TODO: add "check stuff edit rights" middleware
			{
				// POST /edit/stuff/participant adds participant
				stuff.POST("/participant", participantControllers.AddParticipant)
				// POST /edit/stuff/judge adds judge
				stuff.POST("/judge", judgeControllers.AddJudge)

				// DELETE /edit/stuff/participant deletes participant
				stuff.DELETE("/participant/:id", participantControllers.DeleteParticipant)
				// DELETE /edit/stuff/judge deletes judge
				stuff.DELETE("/judge/:id", judgeControllers.DeleteJudge)

				// PATCH /edit/stuff/participant updates participant
				stuff.PATCH("/participant", participantControllers.UpdateParticipant)
				// PATCH /edit/stuff/judge updates judge
				stuff.PATCH("/judge", judgeControllers.UpdateJudge)
			}

			competition := edit.Group("/competitions")
			competition.Use() // TODO: add "check competition edit rights" middleware
			{
				// DELETE /edit/competitions/:id deletes competition
				competition.DELETE("/:id")
				// PATCH /edit/competitions/:id updates competition
				competition.PATCH("/")
			}
		}

		// Allowed for all authorized users
		// POST /competitions adds competition
		authorized.POST("/competitions", competitionsControllers.NewCompetition)
		// GET /my-competition provides your own(or to which you were invited) competitions list
		authorized.GET("/my-competitions", competitionsControllers.GetMyCompetitionsShort)
	}

	go healthCheck(log)

	log.Infof("Server launched and running on http://localhost:%s\n", cfg.ServerDevPort)
	log.Fatal(r.Run(":" + cfg.ServerDevPort))
}

func healthCheck(log *zap.SugaredLogger) {
	for {
		time.Sleep(time.Minute)
		log.Info("Health check: I`m fine :)")
	}
}
