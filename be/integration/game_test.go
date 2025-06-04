package integration

import (
	"context"
	"testing"
    "fmt"

	"github.com/noahsignt/blackout/be/model"
	"github.com/noahsignt/blackout/be/repository"
	"github.com/noahsignt/blackout/be/service"
    "github.com/noahsignt/blackout/be/errors"

	"github.com/stretchr/testify/require"
)

func createTestGame(t *testing.T) (context.Context, *service.GameService, *service.PlayerService, *model.Game) {
	ctx := context.Background()

	db := client.Database("testdb")
	gameRepo := repository.NewGameRepo(db)
    playerRepo := repository.NewPlayerRepo(db)

    playerService := service.NewPlayerService(playerRepo)
	gameService := service.NewGameService(gameRepo, playerService)

	game := &model.Game{}
	createdGame, err := gameService.CreateGame(ctx, game)
	require.NoError(t, err)

	return ctx, gameService, playerService, createdGame
}

func TestFindGame(t *testing.T) {
	ctx, service, _, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	found, err := service.GetGameByID(ctx, game.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
}

func TestJoinGame(t *testing.T) {
    ctx, gameService, playerService, game := createTestGame(t)

    t.Logf("Created game: %+v", game)

    player, err := playerService.CreatePlayer(ctx, &model.Player{Name: "TestJoinGamePlayer"})
    require.NoError(t, err)
    updatedGame, err := gameService.JoinGame(ctx, game.ID, player.ID)

    require.NoError(t, err)
    require.NotNil(t, updatedGame)

    foundPlayer := updatedGame.Players[0]
    require.Equal(t, *player, foundPlayer)
}

func TestStartGame0Players(t *testing.T) {
	ctx, service, _, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	_, err := service.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooFewPlayers)
}

func TestStartGame7Players(t *testing.T) {
	ctx, gameService, playerService, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

    for i := range 7 {
        playerIDHex := fmt.Sprintf("player-%d", i)
        player, err := playerService.CreatePlayer(ctx, &model.Player{Name: playerIDHex})
        require.NoError(t, err)
        _, err = gameService.JoinGame(ctx, game.ID, player.ID)
        if err != nil {
            t.Fatalf("failed to join game with player %d: %v", i, err)
        }
    }

	_, err := gameService.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooManyPlayers)
}


func TestStartGame3Players(t *testing.T) {
    ctx, gameService, playerService, game := createTestGame(t)

    t.Logf("Created game: %+v", game)

    for i := range 3 {
        playerIDHex := fmt.Sprintf("player-%d", i)
        player, err := playerService.CreatePlayer(ctx, &model.Player{Name: playerIDHex})
        require.NoError(t, err)
        _, err = gameService.JoinGame(ctx, game.ID, player.ID)
        if err != nil {
            t.Fatalf("failed to join game with player %d: %v", i, err)
        }
    }

    updatedGame, err := gameService.StartGame(ctx, game.ID)
    round := updatedGame.Round
    hand := round.CurrHand

    t.Logf("Updated game: %+v", updatedGame)

    require.Equal(t, 1, round.RoundNum)
    require.GreaterOrEqual(t, round.Trump, 1)
    require.LessOrEqual(t, round.Trump, 4)
    require.Equal(t, 0, hand.StartingIdx)
    require.NoError(t, err)
}