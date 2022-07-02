package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type IRepo interface {
	NewCompetition(collection *mongo.Collection, competition models.Competition) error
	GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(collection *mongo.Collection) (models.Competition, error)
}
