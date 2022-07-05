package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"uacs/internal/models"
)

type ICompetitionRepoV0 interface {
	CreateCompetition(collection *mongo.Collection, competition models.Competition) error
	DeleteCompetition(collection *mongo.Collection, id string) error
	CreateCompetitionParticipantsEntity(collection *mongo.Collection, entity models.CompetitionParticipantsEntity) error
	CreateCompetitionQualificationEntity(collection *mongo.Collection, entity models.CompetitionQualificationEntity) error
	UpdateCompetition(collection *mongo.Collection, competition models.Competition) error

	GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error)
	GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error)
	GetSingleCompetitionFull(collection *mongo.Collection, id string) (models.Competition, error)
}

type IParticipantRepoV0 interface {
	AddParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error
	DeleteParticipant(collection *mongo.Collection, id string) error
	UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error
	GetParticipant(collection *mongo.Collection, id, competitionId, division, ac string) (models.CompetitionParticipant, error)
	GetParticipants(collection *mongo.Collection, competitionId string) ([]models.CompetitionParticipantShortOutput, error)
}

type IJudgeRepoV0 interface {
	AddJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error
	DeleteJudge(collection *mongo.Collection, id string) error
	UpdateJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error
	GetJudge(collection *mongo.Collection, id string, competitionId string) (models.CompetitionJudge, error)
	GetJudges(collection *mongo.Collection, competitionId string) ([]models.CompetitionJudge, error)
}
