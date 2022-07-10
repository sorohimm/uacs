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

type ParticipantsRepoV0 struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func (r *ParticipantsRepoV0) CreateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	filter := bson.M{"competitionUUID": competitionId}
	update := bson.D{{"$push", bson.D{{"participants", participant}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantsRepoV0) GetParticipant(collection *mongo.Collection, competitionId string, id string) (models.CompetitionParticipant, error) {
	filter := bson.D{
		{"competitionUUID", competitionId},
		{"participants", bson.M{"$elemMatch": bson.M{"uuid": id}}},
	}

	projection := bson.M{"_id": 0}
	opt := options.FindOne().SetProjection(projection)

	var participant models.CompetitionParticipantsEntity
	err := collection.FindOne(context.Background(), filter, opt).Decode(&participant)
	if err != nil {
		return models.CompetitionParticipant{}, err
	}

	result := getParticipant(participant, id)
	if result == nil {
		return models.CompetitionParticipant{}, mongo.ErrNoDocuments
	}

	return *result, nil
}

func getParticipant(participants models.CompetitionParticipantsEntity, id string) *models.CompetitionParticipant {
	for _, el := range participants.Participants {
		if el.UUID == id {
			return &el
		}
	}
	return nil
}

func (r *ParticipantsRepoV0) GetParticipants(collection *mongo.Collection, competitionId string) (models.CompetitionParticipantsEntity, error) {
	var result models.CompetitionParticipantsEntity
	err := collection.FindOne(context.Background(), bson.M{"competitionUUID": competitionId}).Decode(&result)
	if err != nil {
		return models.CompetitionParticipantsEntity{}, err
	}

	return result, nil
}

func (r *ParticipantsRepoV0) GetDivisionParticipants(collection *mongo.Collection, competitionId, division string) (models.CompetitionParticipantsEntity, error) {
	filter := bson.D{
		{"competitionUUID", competitionId},
		{"participants", bson.M{"$elemMatch": bson.M{"uuid": division}}},
	}

	projections := bson.M{
		"_id": 0,
	}

	opt := options.FindOne().SetProjection(projections)

	var participant models.CompetitionParticipantsEntity
	err := collection.FindOne(context.Background(), filter, opt).Decode(&participant)
	if err != nil {
		return models.CompetitionParticipantsEntity{}, err
	}
	return participant, nil
}

func (r *ParticipantsRepoV0) DeleteParticipant(collection *mongo.Collection, competitionId, id string) error {
	filter := bson.M{"competitionUUID": competitionId}
	update := bson.D{{"$pull", bson.M{"participants": bson.M{"uuid": id}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantsRepoV0) UpdateParticipant(collection *mongo.Collection, participant models.CompetitionParticipant, competitionId string) error {
	return nil
}
