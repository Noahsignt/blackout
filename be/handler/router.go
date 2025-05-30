package handler

import (
    "github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
    r := chi.NewRouter()

    gameHandler := NewGameHandler()

    // Set up routes
    r.Get("/game/{id}", gameHandler.GetGameByID)
    r.Post("/game", gameHandler.CreateGame)

    return r
}
