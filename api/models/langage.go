package models

import objectid "go.mongodb.org/mongo-driver/bson/primitive"

// Language is db model to register mongodb
type Language struct {
	ID       objectid.ObjectID `bson: "_id,omitempty"`
	Japanese string            `bson:"japan"`
	English  string            `bson:"English"`
}
