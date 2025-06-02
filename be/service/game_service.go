package service

import (
	"context"
	"errors"
    "fmt"

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
    }, nil
}

func (s *GameService) CreateGame(ctx context.Context, game *model.Game) (*model.Game, error) {
    if game == nil {
        return nil, errors.New("game is nil")
    }

    createdGame, err := s.gameRepo.CreateGame(ctx, *game)
    if err != nil {
        return nil, fmt.Errorf("failed to create game: %w", err)
    }

    return createdGame, nil
}

