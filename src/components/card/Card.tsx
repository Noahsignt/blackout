import type { CardType } from "../../game/Game";
import type { ReactElement } from 'react';
import { heart, club, spade, diamond } from "./symbols";

const suitSymbols: Record<CardType["suit"], { icon: ReactElement; }> = {
    s: { icon: spade},
    h: { icon: heart },
    d: { icon: diamond },
    c: { icon: club },
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
    const colorClass = suit === 'h' || suit === 'd' ? 'text-red-600' : 'text-black';
  
    return (
      <div
        className={`border border-gray-300 rounded-lg 
          w-20 h-30 p-2
          bg-white
          flex flex-col justify-between items-center
          text-2xl font-bold
          ${colorClass}`}
      >
        <div>{value}</div>
        <div>{suitData.icon}</div>
        <div className="rotate-180">{value}</div>
      </div>
    );
  }
  
