package model

import (
    "go.mongodb.org/mongo-driver/v2/bson"
)

type Player struct {
    ID     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	GameID bson.ObjectID `bson:"game_id" json:"game_id"`
    Score  int                `bson:"score" json:"score"`
    UserID bson.ObjectID `bson:"user_id" json:"user_id"`
}
