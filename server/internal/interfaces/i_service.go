package interfaces

import "uacs/internal/models"

type IServicesV0 interface {
	NewCompetition(competition models.Competition) (models.Competition, error)

	GetMyCompetitionsShort(userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(id string) (models.Competition, error)

	AddParticipant(participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
	AddJudge(judge models.CompetitionJudge) (models.CompetitionJudge, error)

	DeleteParticipant(id string) error
	DeleteJudge(id string) error

	UpdateParticipant(participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
	UpdateJudge(judge models.CompetitionJudge) (models.CompetitionJudge, error)

	DeleteCompetition(id string) error
	UpdateCompetition(competition models.Competition) (models.Competition, error)

	GetJudges(competitionId string) (models.CompetitionJudge, error)
	GetJudge(judgeId string) ([]models.CompetitionJudge, error)
	GetParticipants(competitionId string) ([]models.CompetitionParticipantShortOutput, error)
	GetParticipant(participantId string) (models.CompetitionParticipant, error)
}
