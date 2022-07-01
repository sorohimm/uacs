package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type IRepo interface {
	NewCompetition(db *mongo.Database, competition models.Competition) error
}
