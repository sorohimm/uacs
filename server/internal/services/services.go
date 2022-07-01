package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/interfaces"
	"uacs/internal/models"
)

type Services struct {
	Log       *zap.SugaredLogger
	Config    *config.Config
	Repo      interfaces.IRepo
	DbHandler interfaces.IDBHandler
}

func (s *Services) NewCompetition(newCompetition models.Competition) error {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	err := s.Repo.NewCompetition(database, newCompetition)
	if err != nil {
		s.Log.Errorf("Failed create new competition. Received error: %s", err.Error())
		return err
	}
	return nil
}
