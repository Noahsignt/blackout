package service

import (
    "errors"
    "github.com/noahsignt/blackout/be/model"
)

type GameService struct {
    // You can add a repository field here later, e.g.
    // Repo *repository.GameRepository
}

func NewGameService() *GameService {
    return &GameService{}
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
