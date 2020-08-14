package models

import (
	"sync"

	"github.com/ono5/study-hacker/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var instance *mongoDB

// GetDBInstance c
func GetDBInstance() *mongoDB {
	once.Do(func() {
		db := newMongoDB()
		instance = db
	})
	return instance
}

// mongoDB is a struct to connect mongoDB Database
type mongoDB struct {
	collection *mongo.Collection
	client     *mongo.Client
}

// newMongoDB is a constractor
func newMongoDB() *mongoDB {
	m := mongoDB{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	utils.LogFatal("Fatal to create mongo client: %v\n", err)
	m.client = client
	return &m
}

func (m *mongoDB) ConnectDB(dbName, collectionName string) {
	m.collection = m.client.Database(dbName).Collection(collectionName)
}
