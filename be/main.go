package main

import (
    "log"
    "net/http"
    "time"

    "github.com/noahsignt/blackout/be/handler"
    "github.com/noahsignt/blackout/be/config"
    "github.com/noahsignt/blackout/be/repository"
    "github.com/noahsignt/blackout/be/service"
)

func main() {
    // -- Environment Variables / Context --
    ctx := config.Load()

    // -- Repositories --
    gameRepo, playerRepo, userRepo, err := repository.InitRepos(ctx.DBUri, "blackout")
    if err != nil {
        log.Fatal(err)
    }

    // -- Services --
    playerService := service.NewPlayerService(playerRepo)
    gameService := service.NewGameService(gameRepo, playerService)
    userService := service.NewUserService(userRepo, ctx.JWTSecret)

    // -- Router --
    router := handler.NewRouter(*gameService, *userService)
    log.Printf("✅ Server started successfully at %s", time.Now().Format(time.RFC3339))
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
