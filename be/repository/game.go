package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type GameRepo struct {
    collection *mongo.Collection
}

func NewGameRepo(db *mongo.Database) *GameRepo {
    return &GameRepo{
        collection: db.Collection("games"),
    }
}
