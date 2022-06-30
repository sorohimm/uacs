package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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

	controllers := injector.InjectController()

	r := gin.Default()
	v0 := r.Group("/v0")
	{
		v0.POST("/new_competition", controllers.CreateUser)
	}

	log.Infof("Server launched and running on http://localhost:%s\n", cfg.Port)
	log.Fatal(r.Run(cfg.Port))
}
