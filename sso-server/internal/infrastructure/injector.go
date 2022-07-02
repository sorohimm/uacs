package infrastructure

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"go.uber.org/zap"
	"uacs/sso-server/internal/config"
	"uacs/sso-server/internal/controllers"
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
		Log:            e.logger,
		KeyloackClient: e.keycloakClient,
		Cfg:            e.config,
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
