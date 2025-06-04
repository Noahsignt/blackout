package integration

import (
    "context"
    "testing"

    "github.com/stretchr/testify/require"
	"github.com/noahsignt/blackout/be/service"
	"github.com/noahsignt/blackout/be/repository"
	"github.com/noahsignt/blackout/be/model"
)


func TestCreateGame(t *testing.T) {
    ctx := context.Background()

    // reuse shared client
    db := client.Database("testdb")
    repo := repository.NewGameRepo(db)
    service := service.NewGameService(*repo)

    game := &model.Game{}
    game, err := service.CreateGame(ctx, game)
    require.NoError(t, err)

    t.Logf("Created game: %+v", game)

    found, err := service.GetGameByID(ctx, game.ID.Hex())
    require.NoError(t, err)
    require.NotNil(t, found)
}