package driver

import "github.com/ono5/study-hacker/api/models"

// ConnectDB is a driver to connect MongoDB
func ConnectDB() *models.MongoDB {
	return models.GetDBInstance()
}
