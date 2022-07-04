package main

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/infrastructure"
)

var (
	log *zap.SugaredLogger
	cfg *config.Config
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()
	cfg, err = config.New()
	if err != nil {
		log.Fatalf("load config error: %s", err.Error())
	}
}

func main() {
	keycloakClient := gocloak.NewClient(cfg.KeycloakURL)

	injector, err := infrastructure.Injector(log, keycloakClient, cfg, context.Background())
	if err != nil {
		log.Fatalf("injection failed: %s", err.Error())
	}

	controllers := injector.InjectControllers()

	r := gin.Default()

	r.POST("/registration", controllers.Registration)
	r.POST("/login", controllers.Login)
	r.GET("/is_valid_token", controllers.ValidateAccessToken)
	r.GET("/user_id", controllers.GetUserId)

	log.Infof("Sso server launched and running on http://localhost:%s", cfg.SsoServerDevPort)

	log.Fatal(r.Run(":" + cfg.SsoServerDevPort))
}
