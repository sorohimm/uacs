package infrastructure

import (
	"context"
	"github.com/pkg/errors"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"uacs/internal/config"
)

type MongoClient struct {
	Client *mongo.Client
}

func InitMongoClient(logger *zap.SugaredLogger, cfg *config.Config, ctx context.Context) (IDBHandler, error) {
	clientOptions := options.Client().ApplyURI(cfg.DBAuthData.URL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err.Error())
		return &MongoClient{}, errors.Wrap(err, "mongo initialization err")
	}

	logger.Info("db client init ok")

	return &MongoClient{Client: client}, nil
}

func (c *MongoClient) AcquireDatabase(name string) *mongo.Database {
	return c.Client.Database(name)
}

func (c *MongoClient) AcquireClient() *mongo.Client {
	return c.Client
}

func (c *MongoClient) AcquireSession() (mongo.Session, error) {
	return c.Client.StartSession()
}

func (c *MongoClient) AcquireCollection(dbName, collName string) *mongo.Collection {
	db := c.AcquireDatabase(dbName)
	return db.Collection(collName)
}
