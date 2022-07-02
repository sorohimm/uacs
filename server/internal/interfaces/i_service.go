package interfaces

import "uacs/internal/models"

type IServices interface {
	NewCompetition(competition models.Competition) (models.Competition, error)
	GetMyCompetitionsShort() ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort() ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull() (models.Competition, error)
}
