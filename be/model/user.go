package model

import (
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
    ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Username     string             `bson:"username" json:"username"`
    PasswordHash string             `bson:"password_hash"`
    ImageURL     string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
    CreatedAt    time.Time          `bson:"created_at"`
}
