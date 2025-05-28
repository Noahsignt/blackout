import { CardType, GameState } from "./Game";

const suits = ["s", "d", "c", "h"] as const;
const nums = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14] as const;

export function buildDeck(): CardType[] {
  const deck: CardType[] = [];

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

export function getWinner(playerCards: GameState["playerCards"]): string | null {
  const suitRank: Record<CardType["suit"], number> = {
    s: 4,
    h: 3,
    d: 2,
    c: 1,
  };

  let winner: string | null = null;
  let bestCard: CardType | null = null;

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
