package model

import (
    "go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
    ID   bson.ObjectID `bson:"_id,omitempty" json:"id"`
    NumRounds int `bson:"numRounds" json:"numRounds"`
    Round Round `bson:"round" json:"round"`
    Players []Player `bson:"players" json:"players"`
}