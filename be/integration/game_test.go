package integration

import (
	"context"
	"testing"

	"github.com/noahsignt/blackout/be/model"
	"github.com/noahsignt/blackout/be/repository"
	"github.com/noahsignt/blackout/be/service"
    "github.com/noahsignt/blackout/be/errors"

	"github.com/stretchr/testify/require"
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

func TestStartGame0Players(t *testing.T) {
	ctx, service, game := createTestGame(t)

	t.Logf("Created game: %+v", game)

	_, err := service.StartGame(ctx, game.ID)
	require.ErrorIs(t, err, errors.ErrTooFewPlayers)
}
