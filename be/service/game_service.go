package service

import (
    "errors"

    "github.com/noahsignt/blackout/be/model"
    "github.com/noahsignt/blackout/be/repository"
)

type GameService struct {
    gameRepo *repository.GameRepo
}

func NewGameService(gameRepo repository.GameRepo) *GameService {
    return &GameService{&gameRepo}
}

func (s *GameService) GetGameByID(id string) (*model.Game, error) {
    if id == "" {
        return nil, errors.New("invalid id")
    }

	// TODO: don't stub games
    return &model.Game{
        ID:   id,
        Name: "Sample Game",
    }, nil
}

func (s *GameService) CreateGame(game *model.Game) (*model.Game, error) {
    if game == nil {
        return nil, errors.New("game is nil")
    }

    game.ID = "generated-id-123"
    return game, nil
}
