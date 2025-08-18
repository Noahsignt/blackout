package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/noahsignt/blackout/be/model"
	"github.com/noahsignt/blackout/be/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GameHandler struct {
    gameService *service.GameService
}

func NewGameHandler(gameService *service.GameService) *GameHandler {
	return &GameHandler{gameService}
}

// DTO for binary IDs
type gameResponse struct {
    ID        string       `json:"id"`
    NumRounds int          `json:"numRounds"`
    Round     model.Round  `json:"round"`
    Players   []model.Player `json:"players"`
}

func gameToResponse(game *model.Game) *gameResponse {
    return &gameResponse{
        ID:        game.ID.Hex(),
        NumRounds: game.NumRounds,
        Round:     game.Round,
        Players:   game.Players,
    }
}

// GET /game/{id}
func (h *GameHandler) GetGameByID(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    idStr := chi.URLParam(r, "id")

    oid, err := bson.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid game ID", http.StatusBadRequest)
        return
    }

    game, err := h.gameService.GetGameByID(ctx, oid)
    if err != nil {
        http.Error(w, "Game not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gameToResponse(game))
}

// POST /game
func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
    var req model.Game
    ctx := r.Context()

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if req.NumRounds == 0 {
        http.Error(w, "Missing required field: number of rounds", http.StatusBadRequest)
        return
    }

    createdGame, err := h.gameService.CreateGame(ctx, &req)
    if err != nil {
        http.Error(w, "Failed to create game", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gameToResponse(createdGame))
}

// POST /game/{id}/start
func (h *GameHandler) StartGame(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    idStr := chi.URLParam(r, "id")

    oid, err := bson.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid game ID", http.StatusBadRequest)
        return
    }

    startedGame, err := h.gameService.StartGame(ctx, oid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gameToResponse(startedGame))
}
