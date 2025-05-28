import type { Game, Move } from "boardgame.io";
import { INVALID_MOVE } from 'boardgame.io/core';

interface Card {
  suit: "s" | "d" | "c" | "h",
  num: 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15
}

export interface GameState {
  cards: Array<Card>,
  playerCards: Record<string, Card>
}

// takes game state G, ctx, playerID
const drawCard: Move<GameState> = ({ G, ctx, playerID }, id) => {   
  const top = G.cards.pop();

  if(!top) {
    return INVALID_MOVE;
  }

  G.playerCards[playerID] = top;
}

const suits = ["s", "d", "c", "h"] as const;
const nums = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15] as const;

function buildDeck(): Card[] {
  const deck: Card[] = [];

  for (const suit of suits) {
    for (const num of nums) {
      deck.push({ suit, num });
    }
  }

  // shuffle the deck
  for (let i = deck.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [deck[i], deck[j]] = [deck[j], deck[i]];
  }

  return deck;
}

function getWinner(playerCards: GameState["playerCards"]): string | null {
  const suitRank: Record<Card["suit"], number> = {
    s: 4,
    h: 3,
    d: 2,
    c: 1,
  };

  let winner: string | null = null;
  let bestCard: Card | null = null;

  for (const [playerID, card] of Object.entries(playerCards)) {
    if (
      !bestCard ||
      card.num > bestCard.num ||
      (card.num === bestCard.num && suitRank[card.suit] > suitRank[bestCard.suit])
    ) {
      bestCard = card;
      winner = playerID;
    }
  }

  return winner;
}

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