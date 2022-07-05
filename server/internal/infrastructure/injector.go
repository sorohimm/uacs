package infrastructure

import (
	"context"
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
	InjectCompetitionsControllers() controllers.CompetitionControllers
	InjectParticipantControllers() controllers.ParticipantControllers
	InjectJudgeControllers() controllers.JudgeControllers
	InjectMiddlewareV0() middleware.MiddlewareV0
}

var env *environment

type environment struct {
	logger   *zap.SugaredLogger
	cfg      *config.Config
	client   *http.Client
	dbClient interfaces.IDBHandler
}

func (e *environment) InjectCompetitionsControllers() controllers.CompetitionControllers {
	return controllers.CompetitionControllers{
		Log: e.logger,
		CompetitionServices: &services.CompetitionServicesV0{
			Log:    e.logger,
			Config: e.cfg,
			CompetitionsRepoV0: &repo.CompetitionsRepoV0{
				Log:    e.logger,
				Config: e.cfg,
			},
			DbHandler: e.dbClient,
		},
	}
}

func (e *environment) InjectParticipantControllers() controllers.ParticipantControllers {
	return controllers.ParticipantControllers{
		Log: e.logger,
		ParticipantServices: &services.ParticipantServicesV0{
			Log:    e.logger,
			Config: e.cfg,
			ParticipantRepoV0: &repo.ParticipantsRepoV0{
				Log:    e.logger,
				Config: e.cfg,
			},
			DbHandler: e.dbClient,
		},
	}
}

func (e *environment) InjectJudgeControllers() controllers.JudgeControllers {
	return controllers.JudgeControllers{
		Log: e.logger,
		JudgeServices: &services.JudgeServicesV0{
			Log:    e.logger,
			Config: e.cfg,
			JudgeRepoV0: &repo.JudgesRepoV0{
				Log:    e.logger,
				Config: e.cfg,
			},
			DbHandler: e.dbClient,
		},
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
