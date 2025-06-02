package main

import (
    "log"
    "net/http"

    "github.com/noahsignt/blackout/be/handler"
    "github.com/noahsignt/blackout/be/config"
    "github.com/noahsignt/blackout/be/repository"
    "github.com/noahsignt/blackout/be/service"
)

func main() {
    // -- Environment Variables / Context --
    ctx := config.Load()

    // -- Game Repository --
    gameRepo, err := repository.InitGameRepo(ctx.DBUri, "blackout")
    if err == nil {
        log.Fatal(err)
    }

    // -- Game Services -> handles logic related to game construction, etc --
    gameService := service.NewGameService(*gameRepo)

    // -- Router --
    router := handler.NewRouter(*gameService)
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
