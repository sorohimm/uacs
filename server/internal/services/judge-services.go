package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/interfaces"
	"uacs/internal/models"
)

type JudgeServicesV0 struct {
	JudgeRepoV0 interfaces.IJudgeRepoV0
	Log         *zap.SugaredLogger
	Config      *config.Config
	DbHandler   interfaces.IDBHandler
}

func (s *JudgeServicesV0) CreateJudge(competitionId string, judge models.CompetitionJudge) (models.CompetitionJudge, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Judges)

	judge.GenerateUUID()

	err := s.JudgeRepoV0.CreateJudge(collection, judge, "")
	if err != nil {
		s.Log.Errorf("Failed to add judge. Received error: %s", err.Error())
		return models.CompetitionJudge{}, err
	}

	return models.CompetitionJudge{}, nil
}

func (s *JudgeServicesV0) DeleteJudge(competitionId string, id string) error {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Judges)

	err := s.JudgeRepoV0.DeleteJudge(collection, id)
	if err != nil {
		s.Log.Errorf("Failed to delete judge. Received error: %s", err.Error())
		return err
	}

	return nil
}

func (s *JudgeServicesV0) UpdateJudge(competitionId string, judge models.CompetitionJudge) (models.CompetitionJudge, error) {
	return models.CompetitionJudge{}, nil
}

func (s *JudgeServicesV0) GetJudges(competitionId string) (models.CompetitionJudge, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	_ = database.Collection(s.Config.Collections.Judges)

	return models.CompetitionJudge{}, nil
}

func (s *JudgeServicesV0) GetJudge(competitionId string, judgeId string) ([]models.CompetitionJudge, error) {
	return nil, nil
}
