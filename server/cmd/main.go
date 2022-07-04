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
		add := authorized.Group("/add")
		{
			add.POST("/participant", controllersV0.AddParticipant)
			add.POST("/judge", controllersV0.AddJudge)
		}

		del := authorized.Group("/delete")
		{
			del.DELETE("/participant", controllersV0.DeleteParticipant)
			del.DELETE("/judge", controllersV0.DeleteJudge)
		}

		upd := authorized.Group("/update")
		{
			upd.PATCH("/participant", controllersV0.UpdateParticipant)
			upd.PATCH("/judge", controllersV0.UpdateJudge)
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
