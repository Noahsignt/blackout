package service

import (
    "context"
    "fmt"

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

func (s *PlayerService) CreatePlayer(ctx context.Context, player *model.Player) (*model.Player, error) {
    if player == nil {
        return nil, fmt.Errorf("player is nil")
    }

    return s.playerRepo.CreatePlayer(ctx, player)
}

func (s *PlayerService) GetPlayerByID(ctx context.Context, id bson.ObjectID) (*model.Player, error) {
    return s.playerRepo.GetPlayerByID(ctx, id)
}

func (s *PlayerService) UpdatePlayer(ctx context.Context, id bson.ObjectID, player *model.Player) (*model.Player, error) {
    return s.playerRepo.UpdatePlayer(ctx, id, *player)
}
