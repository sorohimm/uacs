package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type JudgesRepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *JudgesRepoV0) AddJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error {
	filter := bson.M{"competition_uuid": competitionId}
	update := bson.D{{"$push", bson.D{{"judging_staff", judge}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *JudgesRepoV0) DeleteJudge(collection *mongo.Collection, id string) error {
	return nil
}

func (r *JudgesRepoV0) UpdateJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error {
	return nil
}

func (r *JudgesRepoV0) GetJudge(collection *mongo.Collection, id string, competitionId string) (models.CompetitionJudge, error) {
	p := mongo.Pipeline{
		{{"$match", bson.M{"competition_uuid": competitionId}}},
		{{"judging_staff", bson.M{"$elemMatch": bson.M{"uuid": id}}}},
	}

	var judge models.CompetitionJudge
	err := collection.FindOne(context.Background(), p).Decode(&judge)
	if err != nil {
		return models.CompetitionJudge{}, err
	}
	return judge, nil
}

func (r *JudgesRepoV0) GetJudges(collection *mongo.Collection, competitionId string) ([]models.CompetitionJudge, error) {
	return nil, nil
}
