package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type IRepoV0 interface {
	NewCompetition(collection *mongo.Collection, competition models.Competition) error

	GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(collection *mongo.Collection, id string) (models.Competition, error)

	AddParticipant(collection *mongo.Collection, participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
	AddJudge(collection *mongo.Collection, judge models.CompetitionJudge) (models.CompetitionJudge, error)

	DeleteParticipant(collection *mongo.Collection, id string) error
	DeleteJudge(collection *mongo.Collection, id string) error

	UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant) (models.CompetitionParticipant, error)
	UpdateJudge(collection *mongo.Collection, judge models.CompetitionJudge) (models.CompetitionJudge, error)
}
