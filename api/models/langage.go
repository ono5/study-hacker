package models

import objectid "go.mongodb.org/mongo-driver/bson/primitive"

// Language is db model to register mongodb
type Language struct {
	ID       objectid.ObjectID `json:"id" bson:"_id,omitempty"`
	Japanese string            `json:"japanese" bson:"japanese"`
	English  string            `json:"english" bson:"English"`
}
