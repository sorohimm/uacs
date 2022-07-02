package interfaces

import "uacs/internal/models"

type IServicesV0 interface {
	NewCompetition(competition models.Competition) (models.Competition, error)
	GetMyCompetitionsShort(userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull() (models.Competition, error)
}
