package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
)

type Services struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}
