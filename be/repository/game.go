package repository

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"

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

func (r *GameRepo) CreateGame(ctx context.Context, game *model.Game) (*model.Game, error) {
    result, err := r.collection.InsertOne(ctx, game)
    if err != nil {
        return nil, err
    }

    oid, ok := result.InsertedID.(bson.ObjectID)
    if !ok {
        return nil, fmt.Errorf("failed to assert inserted ID as ObjectID")
    }

    game.ID = oid
    return game, nil
}


func (r *GameRepo) GetGameByID(ctx context.Context, id bson.ObjectID) (*model.Game, error) {
    var game model.Game
    filter := bson.M{"_id": id}

    if err := r.collection.FindOne(ctx, filter).Decode(&game); err != nil {
        return nil, err
    }

    return &game, nil
}

func (r *GameRepo) PutGame(ctx context.Context, id bson.ObjectID, game model.Game) (*model.Game, error) {
    filter := bson.M{"_id": id}
    res, err := r.collection.ReplaceOne(ctx, filter, game)
    if err != nil {
        return nil, fmt.Errorf("failed to replace game: %w", err)
    }

    if res.MatchedCount == 0 {
        return nil, fmt.Errorf("game with id %s not found", id)
    }

    return &game, nil
}
