package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/interfaces"
	"uacs/internal/models"
)

type CompetitionServicesV0 struct {
	CompetitionsRepoV0 interfaces.ICompetitionRepoV0
	Log                *zap.SugaredLogger
	Config             *config.Config
	DbHandler          interfaces.IDBHandler
}

func (s *CompetitionServicesV0) CreateCompetition(newCompetition models.Competition) (models.Competition, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	collection := database.Collection(s.Config.Collections.Competitions)

	newCompetition.GenerateUUID()

	// TODO: переделать на транзакцию!!!!!
	err := s.CompetitionsRepoV0.CreateCompetition(collection, newCompetition)
	if err != nil {
		s.Log.Errorf("Failed create new competition. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	collection = database.Collection(s.Config.Collections.Participants)
	participantsEntity := models.CompetitionParticipantsEntity{
		CompetitionUUID: newCompetition.UUID,
		Participants:    []models.CompetitionParticipant{},
	}
	err = s.CompetitionsRepoV0.CreateCompetitionParticipantsEntity(collection, participantsEntity)
	if err != nil {
		s.Log.Errorf("Failed create competition participants entity. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	collection = database.Collection(s.Config.Collections.Qualifications)
	qualificationEntity := models.CompetitionQualificationEntity{CompetitionUUID: newCompetition.UUID}
	err = s.CompetitionsRepoV0.CreateCompetitionQualificationEntity(collection, qualificationEntity)
	if err != nil {
		s.Log.Errorf("Failed create competition qualifications entity. Received error: %s", err.Error())
		return models.Competition{}, err
	}

	return newCompetition, nil
}

func (s *CompetitionServicesV0) GetMyCompetitionsShort(userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (s *CompetitionServicesV0) GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Competitions)

	competitions, err := s.CompetitionsRepoV0.GetAllCompetitionsShort(collection)
	if err != nil {
		s.Log.Errorf("GetAllCompetitionsShort error: %s", err.Error())
		return nil, err
	}

	return competitions, nil
}

func (s *CompetitionServicesV0) GetSingleCompetitionFull(id string) (models.Competition, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Competitions)

	competition, err := s.CompetitionsRepoV0.GetSingleCompetitionFull(collection, id)
	if err != nil {
		s.Log.Errorf("GetSingleCompetitionFull error: %s", err.Error())
		return models.Competition{}, err
	}

	return competition, nil
}

func (s *CompetitionServicesV0) DeleteCompetition(id string) error {
	return nil
}

func (s *CompetitionServicesV0) UpdateCompetition(competition models.Competition) (models.Competition, error) {
	return models.Competition{}, nil
}
