package models

import objectid "go.mongodb.org/mongo-driver/bson/primitive"

// Langage is db model to register mongodb
type Langage struct {
	ID       objectid.ObjectID `bson: "_id,omitempty"`
	Japanese string            `bson:"japan"`
	English  string            `bson:"English"`
}
