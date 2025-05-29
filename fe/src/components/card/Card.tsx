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
    const value = rankNames[card.num] || card.num;
    const suitData = suitSymbols[card.suit];

    // dynamic styling -> colors for suit, hover effect on top cards
    const colorClass = card.suit === 'h' || card.suit === 'd' ? 'text-red-600' : 'text-black';
    const hoverClass = isTop ? "hover:-translate-y-10 transition-transform duration-200 cursor-pointer" : "";

    return (
      <div
        className={`absolute border border-gray-300 rounded-lg w-40 h-60 p-2 bg-white ${hoverClass}`}
        style={{
          zIndex: idx,
          transition: 'transform 0.5s ease-in-out',
          transformStyle: 'preserve-3d',
        }}
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
      </div>
    );
  }