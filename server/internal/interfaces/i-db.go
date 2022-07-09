package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IDBHandler interface {
	AcquireClient() *mongo.Client
	AcquireDatabase(string) *mongo.Database
	AcquireSession() (mongo.Session, error)
	AcquireCollection(dbName, collName string) *mongo.Collection
}
