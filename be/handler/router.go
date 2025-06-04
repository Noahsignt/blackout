package handler

import (
    "github.com/go-chi/chi/v5"

    "github.com/noahsignt/blackout/be/service"
)

func NewRouter(gameService service.GameService) *chi.Mux {
    r := chi.NewRouter()
    gameHandler := NewGameHandler(gameService)

    // Set up routes
    r.Get("/game/{id}", gameHandler.GetGameByID)
    r.Post("/game", gameHandler.CreateGame)

    return r
}
