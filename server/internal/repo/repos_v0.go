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

func (r *RepoV0) CreateCompetitionParticipantsEntity(collection *mongo.Collection, entity models.CompetitionParticipantsEntity) error {
	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoV0) CreateCompetitionQualificationEntity(collection *mongo.Collection, entity models.CompetitionQualificationEntity) error {
	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoV0) GetMyCompetitionsShort(collection *mongo.Collection, userId string) ([]models.CompetitionShortOutput, error) {
	return nil, nil
}

func (r *RepoV0) GetAllCompetitionsShort(collection *mongo.Collection) ([]models.CompetitionShortOutput, error) {
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

func (r *RepoV0) GetSingleCompetitionFull(collection *mongo.Collection, id string) (models.Competition, error) {
	res := collection.FindOne(context.Background(), bson.M{"uuid": id})

	var competition models.Competition
	err := res.Decode(&competition)
	if err != nil {
		return models.Competition{}, err
	}

	return competition, nil
}

func (r *RepoV0) AddParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	filter := bson.M{"competition_uuid": competitionId}
	update := bson.D{{"$push", bson.D{{participant.Class, participant}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoV0) AddJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error {
	filter := bson.M{"competition_uuid": competitionId}
	update := bson.D{{"$push", bson.D{{"judging_staff", judge}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoV0) DeleteParticipant(collection *mongo.Collection, id string) error {
	return nil
}

func (r *RepoV0) DeleteJudge(collection *mongo.Collection, id string) error {
	return nil
}

func (r *RepoV0) UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	return nil
}

func (r *RepoV0) UpdateJudge(collection *mongo.Collection, judge models.CompetitionJudge, competitionId string) error {
	return nil
}

func (r *RepoV0) DeleteCompetition(collection *mongo.Collection, id string) error {
	return nil
}

func (r *RepoV0) UpdateCompetition(collection *mongo.Collection, competition models.Competition) error {
	return nil
}

func (r *RepoV0) GetJudge(collection *mongo.Collection, id string, competitionId string) (models.CompetitionJudge, error) {
	p := mongo.Pipeline{
		{{"$match", bson.M{"competition_uuid": competitionId}}},
		{{"$elemMatch", bson.M{"uuid": id}}},
	}

	var judge models.CompetitionJudge
	err := collection.FindOne(context.Background(), p).Decode(&judge)
	if err != nil {
		return models.CompetitionJudge{}, err
	}
	return judge, nil
}

func (r *RepoV0) GetJudges(collection *mongo.Collection, competitionId string) ([]models.CompetitionJudge, error) {
	return nil, nil
}

func (r *RepoV0) GetParticipant(collection *mongo.Collection, id string, competitionId string) (models.CompetitionParticipant, error) {
	p := mongo.Pipeline{
		{{"$match", bson.M{"competition_uuid": competitionId}}},
		{{"$elemMatch", bson.M{"uuid": id}}},
	}

	var participant models.CompetitionParticipant
	err := collection.FindOne(context.Background(), p).Decode(&participant)
	if err != nil {
		return models.CompetitionParticipant{}, err
	}
	return participant, nil
}

func (r *RepoV0) GetParticipants(collection *mongo.Collection, competitionId string) ([]models.CompetitionParticipantShortOutput, error) {
	projection := bson.M{
		"IRMStatus":         0,
		"code":              0,
		"sex":               0,
		"birthDate":         0,
		"targetNumber":      0,
		"region2":           0,
		"regionName2":       0,
		"region3":           0,
		"regionName3":       0,
		"division":          0,
		"class":             0,
		"discharge":         0,
		"isIndividual":      0,
		"isTeam":            0,
		"isIndividualFinal": 0,
		"isTeamFinal":       0,
		"isMixFinal":        0,
		"isWheelchair":      0,
		"email":             0,
	}

	opt := options.Find().SetProjection(projection)
	p := mongo.Pipeline{
		{{"$match", bson.M{"competition_uuid": competitionId}}},
		{{"$elemMatch", bson.M{"uuid": id}}},
	}
	cursor, err := collection.Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, err
	}

	var participants []models.CompetitionParticipantShortOutput
	for cursor.Next(context.TODO()) {
		var participant models.CompetitionParticipantShortOutput
		if err = cursor.Decode(&participant); err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, nil
}
