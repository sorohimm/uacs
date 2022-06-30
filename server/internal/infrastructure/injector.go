package infrastructure

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/config"
	"uacs/internal/controllers"
	"uacs/internal/interfaces"
	"uacs/internal/services"
)

type IInjector interface {
	InjectController() controllers.Controllers
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	client   *http.Client
	dbClient interfaces.IDBHandler
}

func (e *environment) InjectController() controllers.Controllers {
	return controllers.Controllers{
		Log: e.logger,
		Services: &services.Services{
			Log:    e.logger,
			Config: e.cfg,
		},
	}
}

func Injector(log *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	log.Info("injector starting...")
	client, err := initPostgresClient(log, cfg, ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("db init ok")

	env = &environment{
		logger:   log,
		cfg:      cfg,
		client:   http.DefaultClient,
		dbClient: client,
	}

	log.Info("injecting done")
	return env, nil
}
