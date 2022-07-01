package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type IRepo interface {
	NewCompetition(collection *mongo.Collection, competition models.Competition) error
}
