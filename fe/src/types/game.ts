// DTOs
export interface CreateGameRequest {
    numRounds: number;
}

export interface GameResponse {
    id: string;
    numRounds: number;
    round: {
        number: number;
        hand: {
            turn: number;
        };
    };
    players: Array<{
        id: string;
        name: string;
    }>;
}

export type CreateGameResponse = GameResponse;
export type StartGameResponse = GameResponse;