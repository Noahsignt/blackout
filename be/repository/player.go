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

func (r *PlayerRepo) UpdatePlayerScore(ctx context.Context, id bson.ObjectID, newScore int) (*model.Player, error) {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"score": newScore}}

    res, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return nil, err
    }
    if res.MatchedCount == 0 {
        return nil, fmt.Errorf("player not found")
    }

    return r.GetPlayerByID(ctx, id)
}