// DTOs
export interface CreateGameRequest {
    numRounds: number;
}

export interface Player {
    id: string;
    name: string;
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
    players: Array<Player>;
}

export type CreateGameResponse = GameResponse;
export type StartGameResponse = GameResponse;