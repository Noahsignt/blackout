import type { CardType } from "../../game/Game";
import type { ReactElement } from 'react';
import { heart, club, spade, diamond } from "./symbols";

const suitSymbols: Record<CardType["suit"], { icon: ReactElement; color: string }> = {
    s: { color: 'black', icon: spade},
    h: { color: 'red', icon: heart },
    d: { color: 'red', icon: diamond },
    c: { color: 'black', icon: club },
};

const rankNames: Record<number, string> = {
  11: "J",
  12: "Q",
  13: "K",
  14: "A"
};

export function Card({ suit, num }: CardType) {
  const value = rankNames[num] || num;
  const suitData = suitSymbols[suit];

  return (
    <div style={{
      border: '1px solid #ccc',
      borderRadius: '8px',
      width: '80px',
      height: '120px',
      padding: '8px',
      backgroundColor: 'white',
      color: suitData.color,
      display: 'flex',
      flexDirection: 'column',
      justifyContent: 'space-between',
      alignItems: 'center',
      fontSize: '24px',
      fontWeight: 'bold',
    }}>
      <div>{value}</div>
      <div>{suitData.icon}</div>
      <div style={{ transform: 'rotate(180deg)' }}>{value}</div>
    </div>
  );
}
