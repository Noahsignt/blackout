package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/noahsignt/blackout/be/model"
)

type GameRepo struct {
    collection *mongo.Collection
}

func NewGameRepo(db *mongo.Database) *GameRepo {
    return &GameRepo{
        collection: db.Collection("games"),
    }
}

func (r *GameRepo) CreateGame(ctx context.Context, game model.Game) (*model.Game, error) {
    result, err := r.collection.InsertOne(ctx, game)
    if err != nil {
        return nil, err
    }

    if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
        game.ID = oid.Hex()
    }

    return &game, nil
}
