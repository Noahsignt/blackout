import { authenticatedRequest } from "./util"
import { CreateGameRequest, CreateGameResponse, StartGameResponse, GameResponse } from "../types/game"

export const createGame = (request: CreateGameRequest): Promise<CreateGameResponse> => {
    return authenticatedRequest('/api/game', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request)
    });
}

export const joinGame = (gameId: string): Promise<GameResponse> => {
    return authenticatedRequest(`/api/game/${gameId}/join`, {
        method: 'POST'
    });
}

export const startGame = (gameId: string): Promise<StartGameResponse> => {
    return authenticatedRequest(`/api/game/${gameId}/start`, {
        method: 'POST'
    });
}

export const getGame = (gameId: string): Promise<GameResponse> => {
    return authenticatedRequest(`/api/game/${gameId}`, {
        method: 'GET'
    });
}