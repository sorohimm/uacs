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
}
