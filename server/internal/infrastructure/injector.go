package infrastructure

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"uacs/internal/config"
	"uacs/internal/controllers"
	"uacs/internal/interfaces"
	"uacs/internal/middleware"
	"uacs/internal/repo"
	"uacs/internal/services"
)

type IInjector interface {
	InjectControllersV0() controllers.ControllersV0
	InjectMiddlewareV0() middleware.MiddlewareV0
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	client   *http.Client
	dbClient interfaces.IDBHandler
}

func (e *environment) InjectControllersV0() controllers.ControllersV0 {
	return controllers.ControllersV0{
		Log: e.logger,
		ServicesV0: &services.ServicesV0{
			Log:    e.logger,
			Config: e.cfg,
			RepoV0: &repo.RepoV0{
				Log:    e.logger,
				Config: e.cfg,
			},
			DbHandler: e.dbClient,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectMiddlewareV0() middleware.MiddlewareV0 {
	return middleware.MiddlewareV0{
		Log:        e.logger,
		HttpClient: e.client,
		Cfg:        e.cfg,
	}
}

func Injector(log *zap.SugaredLogger, ctx context.Context, cfg *config.Config) (IInjector, error) {
	client, err := InitMongoClient(log, cfg, ctx)
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

	return env, nil
}
