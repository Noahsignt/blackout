import { authenticatedRequest } from "./util"
import { CreateGameRequest, CreateGameResponse, StartGameResponse } from "../types/game"

export const createGame = (request: CreateGameRequest): Promise<CreateGameResponse> => {
    return authenticatedRequest('/api/game', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request)
    });
}

export const startGame = (gameId: string): Promise<StartGameResponse> => {
    return authenticatedRequest(`/api/game/${gameId}/start`, {
        method: 'POST'
    });
}