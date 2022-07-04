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

	// TODO: переделать на транзакцию!!!!!
	err := s.RepoV0.NewCompetition(collection, newCompetition)
	if err != nil {
		s.Log.Errorf("Failed create new competition. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	collection = database.Collection(s.Config.Collections.Participants)
	participantsEntity := models.CompetitionParticipantsEntity{CompetitionUUID: newCompetition.UUID}
	err = s.RepoV0.CreateCompetitionParticipantsEntity(collection, participantsEntity)
	if err != nil {
		s.Log.Errorf("Failed create competition participants entity. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	collection = database.Collection(s.Config.Collections.Qualifications)
	qualificationEntity := models.CompetitionQualificationEntity{CompetitionUUID: newCompetition.UUID}
	err = s.RepoV0.CreateCompetitionQualificationEntity(collection, qualificationEntity)
	if err != nil {
		s.Log.Errorf("Failed create competition qualifications entity. Received error: %s", err.Error())
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

func (s *ServicesV0) GetSingleCompetitionFull(id string) (models.Competition, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	collection := database.Collection(s.Config.Collections.Competitions)

	competition, err := s.RepoV0.GetSingleCompetitionFull(collection, id)
	if err != nil {
		s.Log.Errorf("GetSingleCompetitionFull error: %s", err.Error())
		return models.Competition{}, err
	}

	return competition, nil
}

func (s *ServicesV0) AddParticipant(participant models.CompetitionParticipant) (models.CompetitionParticipant, error) {
	return models.CompetitionParticipant{}, nil
}

func (s *ServicesV0) AddJudge(judge models.CompetitionJudge) (models.CompetitionJudge, error) {
	return models.CompetitionJudge{}, nil
}

func (s *ServicesV0) DeleteParticipant(id string) error {
	return nil
}

func (s *ServicesV0) DeleteJudge(id string) error {
	return nil
}

func (s *ServicesV0) UpdateParticipant(participant models.CompetitionParticipant) (models.CompetitionParticipant, error) {
	return models.CompetitionParticipant{}, nil
}

func (s *ServicesV0) UpdateJudge(judge models.CompetitionJudge) (models.CompetitionJudge, error) {
	return models.CompetitionJudge{}, nil
}

func (s *ServicesV0) DeleteCompetition(id string) error {
	return nil
}

func (s *ServicesV0) UpdateCompetition(competition models.Competition) (models.Competition, error) {
	return models.Competition{}, nil
}

func (s *ServicesV0) GetJudges(competitionId string) (models.CompetitionJudge, error) {
	return models.CompetitionJudge{}, nil
}

func (s *ServicesV0) GetJudge(judgeId string) ([]models.CompetitionJudge, error) {
	return nil, nil
}

func (s *ServicesV0) GetParticipants(competitionId string) ([]models.CompetitionParticipantShortOutput, error) {
	return nil, nil
}

func (s *ServicesV0) GetParticipant(participantId string) (models.CompetitionParticipant, error) {
	return models.CompetitionParticipant{}, nil
}
