package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/noahsignt/blackout/be/config"
	"github.com/noahsignt/blackout/be/middleware"
	"github.com/noahsignt/blackout/be/service"
)

func NewRouter(ctx config.Config, gameService service.GameService, userService service.UserService) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(middleware.NewCORSMiddleware(ctx.AllowedOrigins))

	gameHandler := NewGameHandler(&gameService, &userService)
	userHandler := NewUserHandler(&userService)

	// unprotected routes
	r.Post("/users/signup", userHandler.SignUp)
	r.Post("/users/login", userHandler.Login)

	// protected
	r.Route("/api", func(api chi.Router) {
		api.Use(userService.AuthMiddleware)

		// game routes
		api.Get("/game/{id}", gameHandler.GetGameByID)
		api.Post("/game", gameHandler.CreateGame)
		api.Post("/game/{id}/join", gameHandler.JoinGame)
		api.Post("/game/{id}/start", gameHandler.StartGame)

		// user routes
		api.Post("/users/{id}/password", userHandler.ChangePassword)
		api.Post("/users/{id}/image", userHandler.UpdateProfileImage)
	})

	return r
}
