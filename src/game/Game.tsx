import type { Game, Move } from "boardgame.io";
import { INVALID_MOVE } from 'boardgame.io/core';

// utility functions for creating game state, handling moves
import { buildDeck, getWinner } from "./util";

export interface CardType {
  suit: "s" | "d" | "c" | "h",
  num: 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14
}

export interface GameState {
  cards: Array<CardType>,
  playerCards: Record<string, CardType>
}

// takes game state G, ctx, playerID
const drawCard: Move<GameState> = ({ G, ctx, playerID }, id) => {   
  const top = G.cards.pop();

  if(!top) {
    return INVALID_MOVE;
  }

  G.playerCards[playerID] = top;
}

// actual game data
export const Blackout : Game<GameState>  = {
  // setup -> sets initital value of game state G which then gets updated as we go
  setup: () => ({ cards: buildDeck(), playerCards: {} }),

  // moves -> moves take the game state and a player id + whatever input you need
  moves: {
    drawCard
  },

  turn: {
    minMoves: 1,
    maxMoves: 1,
  },

  endIf: ({ G, ctx }) => {
    if(ctx.numPlayers === Object.keys(G.playerCards).length) {
      return { winner: getWinner(G.playerCards) }
    }
  }
};