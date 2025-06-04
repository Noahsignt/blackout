package model

import (
    "go.mongodb.org/mongo-driver/v2/bson"
)

type Player struct {
	ID bson.ObjectID			    `bson:"_id,omitempty" json:"id"`
	Score int               `bson:"score" json:"score"`
	Name string    			`bson:"name" json:"name"`
}