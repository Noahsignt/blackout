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

func createTestGame(t *testing.T) (context.Context, *service.GameService, *service.PlayerService, *service.UserService, *model.Game) {
	ctx := context.Background()

	db := client.Database("testdb")
	gameRepo := repository.NewGameRepo(db)
    playerRepo := repository.NewPlayerRepo(db)
    userRepo := repository.NewUserRepo(db)

    playerService := service.NewPlayerService(playerRepo)
	gameService := service.NewGameService(gameRepo, playerService)
    userService := service.NewUserService(userRepo, "automated-testing-secret")

	game := &model.Game{}
	createdGame, err := gameService.CreateGame(ctx, game)
	require.NoError(t, err)

	return ctx, gameService, playerService, userService, createdGame
}

func TestFindGame(t *testing.T) {
    SetupTest(t)
	ctx, service, _, _, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	found, err := service.GetGameByID(ctx, game.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
}

func TestJoinGame(t *testing.T) {
    SetupTest(t)
    ctx, gameService, playerService, userService, game := createTestGame(t)

    t.Logf("Created game: %+v", game)
    
    user, err := userService.SignUp(ctx, "automated_testing_user", "password")
    require.NoError(t, err)

    player, err := playerService.CreatePlayer(ctx, user.ID, game.ID)
    require.NoError(t, err)
    updatedGame, err := gameService.JoinGame(ctx, game.ID, player.ID)

    require.NoError(t, err)
    require.NotNil(t, updatedGame)

    foundPlayer := updatedGame.Players[0]
    require.Equal(t, *player, foundPlayer)
}

func TestStartGame0Players(t *testing.T) {
    SetupTest(t)
	ctx, service, _, _, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	_, err := service.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooFewPlayers)
}

func TestStartGame7Players(t *testing.T) {
    SetupTest(t)
	ctx, gameService, playerService, userService, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

    for i := range 7 {
        username := fmt.Sprintf("automated_testing_user_7playertest_%d", i)
        user, err := userService.SignUp(ctx, username, "password")
        require.NoError(t, err)

        player, err := playerService.CreatePlayer(ctx, user.ID, game.ID)
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
    SetupTest(t)
    ctx, gameService, playerService, userService, game := createTestGame(t)

    t.Logf("Created game: %+v", game)

    for i := range 3 {
        username := fmt.Sprintf("automated_testing_user_3playertest_%d", i)
        user, err := userService.SignUp(ctx, username, "password")
        require.NoError(t, err)

        player, err := playerService.CreatePlayer(ctx, user.ID, game.ID)
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