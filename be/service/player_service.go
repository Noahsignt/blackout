package service

import (
    "context"

    "go.mongodb.org/mongo-driver/v2/bson"
    "github.com/noahsignt/blackout/be/model"
    "github.com/noahsignt/blackout/be/repository"
)

type PlayerService struct {
    playerRepo *repository.PlayerRepo
}

func NewPlayerService(playerRepo *repository.PlayerRepo) *PlayerService {
    return &PlayerService{
        playerRepo: playerRepo,
    }
}

func (s *PlayerService) CreatePlayer(ctx context.Context, userID, gameID bson.ObjectID) (*model.Player, error) {
    player := &model.Player{
        UserID: userID,
        GameID: gameID,
        Score:  0,
    }
    return s.playerRepo.CreatePlayer(ctx, player)
}

func (s *PlayerService) UpdatePlayerScore(ctx context.Context, playerID bson.ObjectID, newScore int) (*model.Player, error) {
    return s.playerRepo.UpdatePlayerScore(ctx, playerID, newScore)
}


func (s *PlayerService) GetPlayerByID(ctx context.Context, id bson.ObjectID) (*model.Player, error) {
    return s.playerRepo.GetPlayerByID(ctx, id)
}