package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/interfaces"
	"uacs/internal/models"
)

type ServicesV0 struct {
	Log       *zap.SugaredLogger
	Config    *config.Config
	RepoV0    interfaces.IRepoV0
	DbHandler interfaces.IDBHandler
}

func (s *ServicesV0) NewCompetition(newCompetition models.Competition) (models.Competition, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	collection := database.Collection(s.Config.Collections.Competitions)

	newCompetition.GenerateUUID()

	err := s.RepoV0.NewCompetition(collection, newCompetition)
	if err != nil {
		s.Log.Errorf("Failed create new competition. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	return newCompetition, nil
}

func (s *ServicesV0) GetMyCompetitionsShort(userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (s *ServicesV0) GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	collection := database.Collection(s.Config.Collections.Competitions)

	competitions, err := s.RepoV0.GetAllCompetitionsShort(collection)
	if err != nil {
		s.Log.Errorf("GetAllCompetitionsShort error: %s", err.Error())
		return nil, err
	}

	return competitions, nil
}

func (s *ServicesV0) GetSingleCompetitionFull() (models.Competition, error) {
	return models.Competition{}, nil
}
