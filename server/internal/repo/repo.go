package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type Repo struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *Repo) NewCompetition(db *mongo.Database, newCompetition models.Competition) error {
	return nil
}
