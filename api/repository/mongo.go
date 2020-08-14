package repository

import (
	"context"
	"fmt"

	"github.com/ono5/study-hacker/api/models"
	"github.com/ono5/study-hacker/api/utils"
	objectid "go.mongodb.org/mongo-driver/bson/primitive"
)

// Create creates Todo Item
func Create(m *models.MongoDB, data *models.Language) (id string, err error) {
	fmt.Println("Create Todo Item")
	res, err := m.Collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("Insert Error", err)
		return "", err
	}
	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		utils.LogFatal("Cannot convert to OID %v\n", err)
	}
	return oid.Hex(), nil
}
