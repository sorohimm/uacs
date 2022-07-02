package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type Repo struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *Repo) NewCompetition(collection *mongo.Collection, newCompetition models.Competition) error {
	_, err := collection.InsertOne(context.Background(), newCompetition)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *Repo) GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *Repo) GetSingleCompetitionFull(collection *mongo.Collection) (models.Competition, error) {
	return models.Competition{}, nil
}
