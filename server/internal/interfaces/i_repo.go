package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type IRepoV0 interface {
	NewCompetition(collection *mongo.Collection, competition models.Competition) error
	CreateCompetitionParticipantsEntity(collection *mongo.Collection, entity models.CompetitionParticipantsEntity) error
	CreateCompetitionQualificationEntity(collection *mongo.Collection, entity models.CompetitionQualificationEntity) error

	GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(collection *mongo.Collection, id string) (models.Competition, error)

	AddParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error
	AddJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error

	DeleteParticipant(collection *mongo.Collection, id string) error
	DeleteJudge(collection *mongo.Collection, id string) error

	UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error
	UpdateJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error

	DeleteCompetition(collection *mongo.Collection, id string) error
	UpdateCompetition(collection *mongo.Collection, competition models.Competition) error

	GetJudge(collection *mongo.Collection, id string, competitionId string) (models.CompetitionJudge, error)
	GetJudges(collection *mongo.Collection, competitionId string) ([]models.CompetitionJudge, error)
	GetParticipant(collection *mongo.Collection, id string, competitionId string) (models.CompetitionParticipant, error)
	GetParticipants(collection *mongo.Collection, competitionId string) ([]models.CompetitionParticipantShortOutput, error)
}
