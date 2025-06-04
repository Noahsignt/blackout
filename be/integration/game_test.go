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
    "go.mongodb.org/mongo-driver/v2/bson"
)

func createTestGame(t *testing.T) (context.Context, *service.GameService, *model.Game) {
	ctx := context.Background()

	db := client.Database("testdb")
	repo := repository.NewGameRepo(db)
	service := service.NewGameService(*repo)

	game := &model.Game{}
	createdGame, err := service.CreateGame(ctx, game)
	require.NoError(t, err)

	return ctx, service, createdGame
}

func TestFindGame(t *testing.T) {
	ctx, service, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	found, err := service.GetGameByID(ctx, game.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
}

func TestJoinGame(t *testing.T) {
    ctx, service, game := createTestGame(t)

    t.Logf("Created game: %+v", game)

    player := model.Player{Name: "TestJoinGamePlayer"}
    updatedGame, err := service.JoinGame(ctx, game.ID, player)
    require.NoError(t, err)
    require.NotNil(t, updatedGame)

    foundPlayer := updatedGame.Players[0]
    require.Equal(t, player, foundPlayer)
}

func TestStartGame0Players(t *testing.T) {
	ctx, service, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	_, err := service.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooFewPlayers)
}

func TestStartGame7Players(t *testing.T) {
	ctx, service, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

    for i := range 7 {
        playerIDHex := fmt.Sprintf("player-%d", i)
        playerID, err := bson.ObjectIDFromHex(playerIDHex)
        if err != nil {
            t.Fatalf("invalid player ID hex string %q: %v", playerIDHex, err)
        }
        _, err = service.JoinGame(ctx, game.ID, model.Player{ID: playerID})
        if err != nil {
            t.Fatalf("failed to join game with player %d: %v", i, err)
        }
    }

	_, err := service.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooManyPlayers)
}


func TestStartGame3Players(t *testing.T) {
    ctx, service, game := createTestGame(t)

    t.Logf("Created game: %+v", game)

    for i := range 3 {
        playerIDHex := fmt.Sprintf("player-%d", i)
        playerID, err := bson.ObjectIDFromHex(playerIDHex)
        if err != nil {
            t.Fatalf("invalid player ID hex string %q: %v", playerIDHex, err)
        }
        _, err = service.JoinGame(ctx, game.ID, model.Player{ID: playerID})
        if err != nil {
            t.Fatalf("failed to join game with player %d: %v", i, err)
        }
    }

    updatedGame, err := service.StartGame(ctx, game.ID)
    round := updatedGame.Round
    hand := round.CurrHand

    t.Logf("Updated game: %+v", updatedGame)

    require.Equal(t, 1, round.RoundNum)
    require.GreaterOrEqual(t, round.Trump, 1)
    require.LessOrEqual(t, round.Trump, 4)
    require.Equal(t, 0, hand.StartingIdx)
    require.NoError(t, err)
}