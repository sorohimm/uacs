package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type CompetitionsRepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *CompetitionsRepoV0) CreateCompetition(collection *mongo.Collection, newCompetition models.Competition) error {
	_, err := collection.InsertOne(context.Background(), newCompetition)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompetitionsRepoV0) CreateCompetitionParticipantsEntity(collection *mongo.Collection, entity models.CompetitionParticipantsEntity) error {
	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompetitionsRepoV0) CreateCompetitionQualificationEntity(collection *mongo.Collection, entity models.CompetitionQualificationEntity) error {
	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompetitionsRepoV0) GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *CompetitionsRepoV0) GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error) {
	projection := bson.M{
		"description":            0,
		"organizedByDescription": 0,
		"competitionRules":       0,
		"tormentType":            0,
		"ageCategories":          0,
		"venue":                  0,
		"timeZone":               0,
	}

	opt := options.Find().SetProjection(projection)

	cursor, err := collection.Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, err
	}

	var competitions []models.CompetitionShortOutput
	for cursor.Next(context.TODO()) {
		var competition models.CompetitionShortOutput
		if err = cursor.Decode(&competition); err != nil {
			return nil, err
		}
		competitions = append(competitions, competition)
	}

	return competitions, nil
}

func (r *CompetitionsRepoV0) GetSingleCompetitionFull(collection *mongo.Collection, id string) (models.Competition, error) {
	res := collection.FindOne(context.Background(), bson.M{"uuid": id})

	var competition models.Competition
	err := res.Decode(&competition)
	if err != nil {
		return models.Competition{}, err
	}

	return competition, nil
}

func (r *CompetitionsRepoV0) DeleteCompetition(collection *mongo.Collection, id string) error {
	return nil
}

func (r *CompetitionsRepoV0) UpdateCompetition(collection *mongo.Collection, competition models.Competition) error {
	return nil
}
