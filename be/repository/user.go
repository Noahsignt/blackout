package repository

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"

    "github.com/noahsignt/blackout/be/model"
)

type UserRepo struct {
    collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
    return &UserRepo{
        collection: db.Collection("users"),
    }
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
    result, err := r.collection.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }

    oid, ok := result.InsertedID.(bson.ObjectID)
    if !ok {
        return nil, fmt.Errorf("failed to assert inserted ID as ObjectID")
    }
    user.ID = oid
    return user, nil
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
    var user model.User
    filter := bson.M{"username": username}
    err := r.collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepo) FindByID(ctx context.Context, userID bson.ObjectID) (*model.User, error) {
    filter := bson.M{"_id": userID}
    res := r.collection.FindOne(ctx, filter)

    var user model.User
    if err := res.Decode(&user); err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("user not found")
        }
        return nil, err
    }

    return &user, nil
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userID bson.ObjectID, newHash string) error {
    filter := bson.M{"_id": userID}
    update := bson.M{"$set": bson.M{"password_hash": newHash}}
    res, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return fmt.Errorf("user not found")
    }
    return nil
}

func (r *UserRepo) UpdateImage(ctx context.Context, userID bson.ObjectID, imageURL string) error {
    filter := bson.M{"_id": userID}
    update := bson.M{"$set": bson.M{"image_url": imageURL}}
    res, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return fmt.Errorf("user not found")
    }
    return nil
}
