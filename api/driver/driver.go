package driver

import "github.com/ono5/study-hacker/api/models"

// ConnectDB is a driver to connect MongoDB
func ConnectDB(dbName, collectionName string) *models.MongoDB {
	db := models.GetDBInstance()
	db.GetCollection(dbName, collectionName)
	return db
}
