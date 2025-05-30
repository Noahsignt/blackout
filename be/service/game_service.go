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

// GetGameByID returns a dummy game or error
func (s *GameService) GetGameByID(id string) (*model.Game, error) {
    // Minimal stub: return a fake game for any id, or error if id empty
    if id == "" {
        return nil, errors.New("invalid id")
    }

    return &model.Game{
        ID:   id,
        Name: "Sample Game",
        // add more fields if needed
    }, nil
}

// CreateGame accepts a game and "creates" it (dummy)
func (s *GameService) CreateGame(game *model.Game) (*model.Game, error) {
    // Just return the same game with a fake ID for now
    if game == nil {
        return nil, errors.New("game is nil")
    }

    game.ID = "generated-id-123"
    return game, nil
}
