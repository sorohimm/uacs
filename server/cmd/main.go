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

	r.GET("/all_competitions", controllersV0.GetAllCompetitionsShort)
	r.GET("/competition", controllersV0.GetSingleCompetitionFull)

	authorized := r.Group("/with_auth")
	authorized.Use(middlewareV0.AuthRequired)
	{
		edit := authorized.Group("/edit")
		{
			stuff := edit.Group("/stuff")
			// TODO: add "check stuff edit rights" middleware
			{
				add := stuff.Group("/add")
				{
					// /with_auth/edit/stuff/add/participant
					add.POST("/participant", controllersV0.AddParticipant)
					// /with_auth/edit/stuff/add/judge
					add.POST("/judge", controllersV0.AddJudge)
				}

				del := stuff.Group("/delete")
				{
					// /with_auth/edit/stuff/delete/participant
					del.DELETE("/participant", controllersV0.DeleteParticipant)
					// /with_auth/edit/stuff/delete/judge
					del.DELETE("/judge", controllersV0.DeleteJudge)
				}

				upd := stuff.Group("/update")
				{
					// /with_auth/edit/stuff/update/participant
					upd.PATCH("/participant", controllersV0.UpdateParticipant)
					// /with_auth/edit/stuff/judge/participant
					upd.PATCH("/judge", controllersV0.UpdateJudge)
				}
			}

			competition := edit.Group("/competition")
			// TODO: add "check competition edit rights" middleware
			{
				// /with_auth/edit/competition/delete/id
				competition.DELETE("/:id")
				// /with_auth/edit/competition/update/id
				competition.PATCH("/:id")
			}
		}

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
