import type { BoardProps } from "boardgame.io/react";
import type { GameState } from "../../game/Game";
import { Card } from "../card/Card";

interface GameProps extends BoardProps<GameState> {}

export function Board(props: GameProps) {
    const { G, ctx, moves } = props;
    const onClick = () => moves.drawCard()

    const cards = G.cards;

    return (
        <div className="flex flex-col justify-end items-center w-full h-[65%]">
          <div className="flex h-80 gap-16 w-full justify-center">
            {Object.keys(G.playerCards).map((playerId, i) => (
            <div className="flex flex-col items-center">
              <div key={i} className="relative w-40 h-60">
                <Card card={G.playerCards[playerId]} idx={1} isTop={false} />
              </div>
              <i>{playerId}</i>
            </div>
            ))}
          </div>
          <div onClick={onClick} className="w-full flex justify-center">
            {cards.map((e, i) => (
              <Card key={i} card={e} idx={i} isTop={i === cards.length - 1} />
            ))}
          </div>
        </div>
      );
      
}