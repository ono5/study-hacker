package repository

import (
	"context"
	"fmt"

	"github.com/ono5/study-hacker/api/models"
	"github.com/ono5/study-hacker/api/utils"
	objectid "go.mongodb.org/mongo-driver/bson/primitive"
)

// CRUD interface defines Create, Read, Update, Delete methods
type CRUD interface {
	Create() (id string, err error)
	Update(data *models.Language) (*models.Language, error)
}

// MongoRepo is a struct
type MongoRepo struct {
	mongoDB *models.MongoDB
	data    *models.Language
}

// NewMongoRepo returns mongorepo data
func NewMongoRepo(mongoDB *models.MongoDB, data *models.Language) *MongoRepo {
	return &MongoRepo{mongoDB: mongoDB, data: data}
}

// Create creates Todo Item
func (mr MongoRepo) Create() (id string, err error) {
	fmt.Println("Create Todo Item")

	res, err := mr.mongoDB.Collection.InsertOne(context.Background(), mr.data)
	if err != nil {
		return "", err
	}
	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		utils.LogFatal("Cannot convert to OID %v\n", err)
	}
	return oid.Hex(), nil
}

// Update updates data related id parameter
func (mr MongoRepo) Update(data *models.Language) (*models.Language, error) {
	fmt.Println("Update Todo Item")

	updateData := &models.Language{}
	filter := objectid.D{{"_id", data.ID}}
	if err := mr.mongoDB.Collection.FindOne(context.Background(), filter).Decode(updateData); err != nil {
		return nil, err
	}

	updateData.ID = data.ID
	updateData.Japanese = data.Japanese
	updateData.English = data.English

	_, updateErr := mr.mongoDB.Collection.ReplaceOne(context.Background(), filter, updateData)
	if updateErr != nil {
		return nil, updateErr
	}

	return updateData, nil
}
