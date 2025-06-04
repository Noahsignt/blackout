package repository

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"

    "github.com/noahsignt/blackout/be/model"
)

type PlayerRepo struct {
    collection *mongo.Collection
}

func NewPlayerRepo(db *mongo.Database) *PlayerRepo {
    return &PlayerRepo{
        collection: db.Collection("players"),
    }
}

func (r *PlayerRepo) CreatePlayer(ctx context.Context, player *model.Player) (*model.Player, error) {
    result, err := r.collection.InsertOne(ctx, player)
    if err != nil {
        return nil, err
    }

    oid, ok := result.InsertedID.(bson.ObjectID)
    if !ok {
        return nil, fmt.Errorf("failed to assert inserted ID as ObjectID")
    }

    player.ID = oid
    return player, nil
}

func (r *PlayerRepo) GetPlayerByID(ctx context.Context, id bson.ObjectID) (*model.Player, error) {
    var player model.Player
    filter := bson.M{"_id": id}

    if err := r.collection.FindOne(ctx, filter).Decode(&player); err != nil {
        return nil, err
    }

    return &player, nil
}

func (r *PlayerRepo) UpdatePlayer(ctx context.Context, id bson.ObjectID, player model.Player) (*model.Player, error) {
    filter := bson.M{"_id": id}
    res, err := r.collection.ReplaceOne(ctx, filter, player)
    if err != nil {
        return nil, fmt.Errorf("failed to update player: %w", err)
    }

    if res.MatchedCount == 0 {
        return nil, fmt.Errorf("player with id %s not found", id)
    }

    return &player, nil
}
