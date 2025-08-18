package service

import (
	"context"
	stdErrors "errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/noahsignt/blackout/be/errors"
	"github.com/noahsignt/blackout/be/model"
	"github.com/noahsignt/blackout/be/repository"
)

type GameService struct {
	gameRepo *repository.GameRepo
    playerService *PlayerService
}

func NewGameService(gameRepo *repository.GameRepo, playerService *PlayerService) *GameService {
    return &GameService{
        gameRepo: gameRepo,
        playerService: playerService,
    }
}

func (s *GameService) GetGameByID(ctx context.Context, id bson.ObjectID) (*model.Game, error) {
	game, err := s.gameRepo.GetGameByID(ctx, id)

	if err != nil {
		return nil, stdErrors.New("could not find game")
	}

	return game, nil
}

func (s *GameService) JoinGame(ctx context.Context, gameID bson.ObjectID, userID bson.ObjectID) (*model.Game, error) {
	game, err := s.gameRepo.GetGameByID(ctx, gameID)

	if err != nil {
		return nil, stdErrors.New("could not find game")
	}

	// check if user is already in the game
	for _, p := range game.Players {
		if p.UserID == userID {
			return nil, fmt.Errorf("user with ID %s is already in this game", userID.Hex())
		}
	}

    // create a new player for this user in this game
    player, err := s.playerService.CreatePlayer(ctx, userID, gameID)

    if err != nil {
        return nil, fmt.Errorf("could not create player: %w", err)
    }

    game.Players = append(game.Players, *player)
    updatedGame, err := s.gameRepo.PutGame(ctx, gameID, *game)

    if err != nil {
        return nil, stdErrors.New("could not update game with new player")
    }

    return updatedGame, nil
}
func (s *GameService) CreateGame(ctx context.Context, game *model.Game) (*model.Game, error) {
	if game == nil {
		return nil, stdErrors.New("game is nil")
	}

	createdGame, err := s.gameRepo.CreateGame(ctx, game)
	if err != nil {
		return nil, fmt.Errorf("failed to create game: %w", err)
	}

	return createdGame, nil
}

func (s *GameService) StartGame(ctx context.Context, id bson.ObjectID) (*model.Game, error) {
	game, err := s.gameRepo.GetGameByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding game: %w", err)
	}

	if game == nil {
		return nil, fmt.Errorf("game with id: %s could not be found", id)
	}

	// validate game is ready to start
	if len(game.Players) < 3 {
		return nil, fmt.Errorf("game with id: %s does not have enough players (%d) to start: %w", id, len(game.Players), errors.ErrTooFewPlayers)
	}

	if len(game.Players) > 6 {
		return nil, fmt.Errorf("game with id: %s has too many players (%d) to start: %w", id, len(game.Players), errors.ErrTooManyPlayers)
	}

	// start a hand with the first player having the first turn
	firstHand := model.NewHand(0)

	// start a round with 1 card being dealt
	firstRound := model.NewRound(1, firstHand)
	game.Round = *firstRound

	updatedGame, err := s.gameRepo.PutGame(ctx, id, *game)
	if err != nil {
		return nil, fmt.Errorf("error updating game: %w", err)
	}

	return updatedGame, nil
}
