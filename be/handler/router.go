package handler

import (
    "github.com/go-chi/chi/v5"

    "github.com/noahsignt/blackout/be/service"
)

func NewRouter(gameService service.GameService, userService service.UserService) *chi.Mux {
    r := chi.NewRouter()
    gameHandler := NewGameHandler(gameService)
    userHandler := NewUserHandler(&userService)

    // Game routes
    r.Get("/game/{id}", gameHandler.GetGameByID)
    r.Post("/game", gameHandler.CreateGame)

    // User routes
    r.Post("/users/signup", userHandler.SignUp)
    r.Post("/users/{id}/password", userHandler.ChangePassword)
    r.Post("/users/{id}/image", userHandler.UpdateProfileImage)

    return r
}