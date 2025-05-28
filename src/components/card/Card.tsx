import type { CardType } from "../../game/Game";
import type { ReactElement } from 'react';
import React, { useState } from 'react';
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

type CardProps = {
    card: CardType,
    idx: number,
    isTop: boolean
}

export function Card({ card, idx, isTop }: CardProps) {
    const [hovered, setHovered] = useState(false);
  
    const baseX = idx * 0.5;
    const baseY = -idx * 0.5;
  
    const value = rankNames[card.num] || card.num;
    const suitData = suitSymbols[card.suit];
    const colorClass = card.suit === 'h' || card.suit === 'd' ? 'text-red-600' : 'text-black';
  
    const transform = hovered && isTop
      ? `translate(${baseX}px, ${baseY - 100}px) rotateY(360deg)`
      : `translate(${baseX}px, ${baseY}px) rotateY(0deg)`;
  
    return (
      <div
        className={`absolute border border-gray-300 rounded-lg w-40 h-60 p-2 bg-white cursor-pointer`}
        style={{
          transform,
          zIndex: idx,
          transition: 'transform 0.5s ease-in-out',
          transformStyle: 'preserve-3d',
        }}
        onMouseEnter={() => setHovered(true)}
        onMouseLeave={() => setHovered(false)}
      >
        {/* front face */}
        <div
          className={`flex flex-col justify-between items-center text-2xl font-bold ${colorClass}`}
          style={{
            backfaceVisibility: 'hidden',
            position: 'absolute',
            width: '100%',
            height: '100%',
            padding: '8px',
            boxSizing: 'border-box',
          }}
        >
          <div>{value}</div>
          <div>{suitData.icon}</div>
          <div className="rotate-180">{value}</div>
        </div>
  
        {/* back face */}
        <div
            className="absolute inset-0 rounded-lg flex items-center justify-center text-white font-bold text-xl"
            style={{
                backfaceVisibility: 'hidden',
                transform: 'rotateY(180deg)',
                backgroundColor: 'black',
                backgroundImage: `
                linear-gradient(45deg, red 25%, transparent 25%),
                linear-gradient(-45deg, red 25%, transparent 25%),
                linear-gradient(45deg, transparent 75%, red 75%),
                linear-gradient(-45deg, transparent 75%, red 75%)
                `,
                backgroundSize: '20px 20px',
                backgroundPosition: '0 0, 0 10px, 10px -10px, -10px 0px',
            }}
            />
      </div>
    );
  }