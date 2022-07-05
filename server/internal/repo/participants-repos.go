package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type ParticipantsRepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *ParticipantsRepoV0) AddParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	filter := bson.M{"competition_uuid": competitionId}
	update := bson.D{{"$push", bson.D{{fmt.Sprintf("%s.%s", participant.Division, participant.AgeClass), participant}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantsRepoV0) GetParticipant(collection *mongo.Collection, id, competitionId, division, ac string) (models.CompetitionParticipant, error) {
	p := mongo.Pipeline{
		{{"$match", bson.M{"competition_uuid": competitionId}}},
		{{fmt.Sprintf("%s.%s", division, ac), bson.M{"$elemMatch": bson.M{"uuid": id}}}},
	}

	var participant models.CompetitionParticipant
	err := collection.FindOne(context.Background(), p).Decode(&participant)
	if err != nil {
		return models.CompetitionParticipant{}, err
	}
	return participant, nil
}

func (r *ParticipantsRepoV0) DeleteParticipant(collection *mongo.Collection, id string) error {
	return nil
}

func (r *ParticipantsRepoV0) UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	return nil
}

func (r *ParticipantsRepoV0) GetParticipants(collection *mongo.Collection, competitionId string) ([]models.CompetitionParticipantShortOutput, error) {
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
	cursor, err := collection.Find(context.Background(), bson.M{"competition_uuid": competitionId}, opt)
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
