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
	injector, err := infrastructure.Injector(log, ctx, cfg)
	if err != nil {
		log.Fatalf("Ijection fatal error: %s\n", err.Error())
	}

	controllersV0 := injector.InjectControllersV0()
	middlewareV0 := injector.InjectMiddlewareV0()
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// public handles
	r.GET("/all_competitions", controllersV0.GetAllCompetitionsShort)
	r.GET("/competition/:id", controllersV0.GetSingleCompetitionFull)

	authorized := r.Group("/")
	authorized.Use(middlewareV0.AuthRequired)
	{
		edit := authorized.Group("/edit")
		{
			stuff := edit.Group("/stuff")
			// TODO: add "check stuff edit rights" middleware
			{
				// POST /edit/stuff/participant adds participant
				stuff.POST("/participant", controllersV0.AddParticipant)
				// POST /edit/stuff/judge adds judge
				stuff.POST("/judge", controllersV0.AddJudge)

				// DELETE /edit/stuff/participant deletes participant
				stuff.DELETE("/participant/:id", controllersV0.DeleteParticipant)
				// DELETE /edit/stuff/judge deletes judge
				stuff.DELETE("/judge/:id", controllersV0.DeleteJudge)

				// PATCH /edit/stuff/participant updates participant
				stuff.PATCH("/participant", controllersV0.UpdateParticipant)
				// PATCH /edit/stuff/judge updates judge
				stuff.PATCH("/judge", controllersV0.UpdateJudge)
			}

			competition := edit.Group("/competitions")
			// TODO: add "check competition edit rights" middleware
			{
				// /edit/competitions/:id
				competition.DELETE("/:id")
				// /edit/competitions/:id
				competition.PATCH("/")
			}
		}

		// allow for all authorized users
		authorized.POST("/new_competition", controllersV0.NewCompetition)
		authorized.GET("/my_competitions", controllersV0.GetMyCompetitionsShort)
	}

	go healthCheck(log)

	log.Infof("Server launched and running on http://localhost:%s\n", cfg.DevPort)
	log.Fatal(r.Run(":" + cfg.DevPort))
}

func healthCheck(log *zap.SugaredLogger) {
	for {
		time.Sleep(time.Minute)
		log.Info("Health check: I`m fine :)")
	}
}
