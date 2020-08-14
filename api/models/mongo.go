package models

import (
	"sync"

	"github.com/ono5/study-hacker/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var instance *MongoDB

// GetDBInstance c
func GetDBInstance() *MongoDB {
	once.Do(func() {
		db := newMongoDB()
		instance = db
	})
	return instance
}

// MongoDB is a struct to connect MongoDB Database
type MongoDB struct {
	collection *mongo.Collection
	client     *mongo.Client
}

// newMongoDB is a constractor
func newMongoDB() *MongoDB {
	m := MongoDB{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	utils.LogFatal("Fatal to create mongo client: %v\n", err)
	m.client = client
	return &m
}

func (m *MongoDB) ConnectDB(dbName, collectionName string) {
	m.collection = m.client.Database(dbName).Collection(collectionName)
}
