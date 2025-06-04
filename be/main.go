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

    // -- Game Repository --
    gameRepo, err := repository.InitGameRepo(ctx.DBUri, "blackout")
    if err != nil {
        log.Fatal(err)
    }

    // -- Game Services -> handles logic related to game construction, etc --
    gameService := service.NewGameService(*gameRepo)

    // -- Router --
    router := handler.NewRouter(*gameService)
    log.Printf("✅ Server started successfully at %s", time.Now().Format(time.RFC3339))
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
