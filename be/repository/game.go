package repository

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *GameRepo) GetGameByID(ctx context.Context, id string) (*model.Game, error) {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("invalid id format: %w", err)
    }

    var game model.Game
    filter := bson.M{"_id": oid}

    if err := r.collection.FindOne(ctx, filter).Decode(&game); err != nil {
        return nil, err
    }

    if game.ID == "" {
        game.ID = oid.Hex()
    }

    return &game, nil
}

func (r *GameRepo) PutGame(ctx context.Context, id string, game model.Game) (*model.Game, error) {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("invalid id format: %w", err)
    }
    game.ID = oid.Hex()

    filter := bson.M{"_id": oid}
    res, err := r.collection.ReplaceOne(ctx, filter, game)
    if err != nil {
        return nil, fmt.Errorf("failed to replace game: %w", err)
    }

    if res.MatchedCount == 0 {
        return nil, fmt.Errorf("game with id %s not found", id)
    }

    return &game, nil
}
