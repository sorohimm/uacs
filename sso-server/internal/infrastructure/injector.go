package infrastructure

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"go.uber.org/zap"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/controllers"
	"uacs/sso-server/internal/services"
)

type IInjector interface {
	InjectControllers() controllers.Controllers
}

var env *environment

type environment struct {
	logger         *zap.SugaredLogger
	keycloakClient gocloak.GoCloak
	config         *config.Config
}

func (e *environment) InjectControllers() controllers.Controllers {
	return controllers.Controllers{
		Log: e.logger,
		Cfg: e.config,
		Services: &services.Services{
			Log:            e.logger,
			Cfg:            e.config,
			KeyloackClient: e.keycloakClient,
		},
	}
}

func Injector(log *zap.SugaredLogger, keycloakClient gocloak.GoCloak, cfg *config.Config, ctx context.Context) (IInjector, error) {
	env = &environment{
		logger:         log,
		keycloakClient: keycloakClient,
		config:         cfg,
	}

	return env, nil
}
