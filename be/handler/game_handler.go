package handler

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/noahsignt/blackout/be/model"
    "github.com/noahsignt/blackout/be/service"
)

type GameHandler struct {
    gameService *service.GameService
}

func NewGameHandler(gameService service.GameService) *GameHandler {
	return &GameHandler{&gameService}
}

// GET /game/{id} -> tries to return the game
func (h *GameHandler) GetGameByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    game, err := h.gameService.GetGameByID(id)
    if err != nil {
        http.Error(w, "Game not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(game)
}

// POST /game -> tries to post the body of request as a new game
func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
    var req model.Game

    ctx := r.Context();

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    game, err := h.gameService.CreateGame(ctx, &req)
    if err != nil {
        http.Error(w, "Failed to create game", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(game)
}
