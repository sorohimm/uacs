package interfaces

import "uacs/internal/models"

type IServices interface {
	NewCompetition(competition models.Competition) error
}
