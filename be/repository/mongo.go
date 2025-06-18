package repository

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitMongo(dbURI string) (*mongo.Client, error) {
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(opts)
    if err != nil {
        return nil, err
    }

    if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to MongoDB!")
    return client, nil
}

func InitRepos(dbURI string, dbName string) (*GameRepo, *PlayerRepo, *UserRepo, error) {
    client, err := InitMongo(dbURI)
    if err != nil {
        return nil, nil, nil, err
    }

    db := client.Database(dbName)
    gameRepo := NewGameRepo(db)
    playerRepo := NewPlayerRepo(db)
    UserRepo := NewUserRepo(db)

    return gameRepo, playerRepo, UserRepo, nil
}
