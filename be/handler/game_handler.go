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
    ctx := r.Context();
    id := chi.URLParam(r, "id")

    game, err := h.gameService.GetGameByID(ctx, id)
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

    // try and decode request body into Game struct
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // validate fields
    if req.NumRounds == 0 {
        http.Error(w, "Missing required field: number of rounds", http.StatusBadRequest)
    }

    game, err := h.gameService.CreateGame(ctx, &req)
    if err != nil {
        http.Error(w, "Failed to create game", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(game)
}
