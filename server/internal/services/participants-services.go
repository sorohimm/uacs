package services

import (
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/infrastructure"
	"uacs/internal/models"
	"uacs/internal/repo"
)

type ParticipantServicesV0 struct {
	ParticipantRepoV0 repo.IParticipantRepoV0
	Log               *zap.SugaredLogger
	Config            *config.Config
	DbHandler         infrastructure.IDBHandler
}

func (s *ParticipantServicesV0) CreateParticipant(competitionId string, participant models.CompetitionParticipant) (models.CompetitionParticipant, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Participants)

	participant.GenerateUUID()

	err := s.ParticipantRepoV0.CreateParticipant(collection, participant, competitionId)
	if err != nil {
		s.Log.Errorf("Failed to add participant. Received error: %s", err.Error())
		return models.CompetitionParticipant{}, err
	}

	return participant, nil
}

func (s *ParticipantServicesV0) UpdateParticipant(competitionId string, participant models.CompetitionParticipant) (models.CompetitionParticipant, error) {
	return models.CompetitionParticipant{}, nil
}

func (s *ParticipantServicesV0) GetParticipants(competitionId string) (models.CompetitionParticipantsEntity, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Participants)

	participants, err := s.ParticipantRepoV0.GetParticipants(collection, competitionId)
	if err != nil {
		s.Log.Error(err)
		return models.CompetitionParticipantsEntity{}, err
	}

	//result := participants.ToShortOutput()

	return participants, nil
}

func (s *ParticipantServicesV0) GetParticipant(competitionId string, id string) (models.CompetitionParticipant, error) {
	collection := s.DbHandler.AcquireCollection(s.Config.DBAuthData.Name, s.Config.Collections.Participants)

	participant, err := s.ParticipantRepoV0.GetParticipant(collection, competitionId, id)
	if err != nil {
		s.Log.Error(err)
		return models.CompetitionParticipant{}, err
	}

	return participant, nil
}

func (s *ParticipantServicesV0) DeleteParticipant(competitionId string, id string) error {
	database := s.DbHandler.AcquireDatabase(s.Config.DBAuthData.Name)
	collection := database.Collection(s.Config.Collections.Participants)

	err := s.ParticipantRepoV0.DeleteParticipant(collection, competitionId, id)
	if err != nil {
		s.Log.Errorf("Failed to delete participant. Received error: %s", err.Error())
		return err
	}

	return nil
}
