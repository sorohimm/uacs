package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"uacs/internal/config"
	"uacs/internal/infrastructure"

	"go.uber.org/zap"
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

	controllersV0 := injector.InjectV0Controllers()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	v0 := r.Group("/v0")
	{
		v0.POST("/new_competition", controllersV0.NewCompetition)
		v0.GET("/my_competitions", controllersV0.GetMyCompetitionsShort)
		v0.GET("/all_competitions", controllersV0.GetAllCompetitionsShort)
		v0.GET("/competition", controllersV0.GetSingleCompetitionFull)
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
