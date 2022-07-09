package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"uacs/internal/config"
	"uacs/internal/models"
)

type ParticipantsRepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *ParticipantsRepoV0) CreateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	filter := bson.M{"competition_uuid": competitionId}
	update := bson.D{{"$push", bson.D{{fmt.Sprintf("%s.%s", participant.Division, participant.AgeCategory), participant}}}}
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

func (r *ParticipantsRepoV0) GetParticipants(collection *mongo.Collection, competitionId string) (models.CompetitionParticipantsEntity, error) {
	var result models.CompetitionParticipantsEntity
	err := collection.FindOne(context.Background(), bson.M{"competition_uuid": competitionId}).Decode(&result)
	if err != nil {
		return models.CompetitionParticipantsEntity{}, err
	}

	return result, nil
}

func (r *ParticipantsRepoV0) DeleteParticipant(collection *mongo.Collection, id string) error {
	return nil
}

func (r *ParticipantsRepoV0) UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	return nil
}
