package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type RepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *RepoV0) NewCompetition(collection *mongo.Collection, newCompetition models.Competition) error {
	_, err := collection.InsertOne(context.Background(), newCompetition)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoV0) GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *RepoV0) GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *RepoV0) GetSingleCompetitionFull(collection *mongo.Collection) (models.Competition, error) {
	return models.Competition{}, nil
}
