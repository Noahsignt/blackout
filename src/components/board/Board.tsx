import type { BoardProps } from "boardgame.io/react";
import type { GameState } from "../../game/Game";
import { Card } from "../card/Card";

interface GameProps extends BoardProps<GameState> {}

export function Board(props: GameProps) {
    const { G, ctx, moves, playerID } = props;
    const onClick = () => moves.drawCard()

    const cards = G.cards;

    return (
        <div className="flex flex-col justify-end items-center w-full h-[65%]">
          <div>
            {cards.map((e, i) => (
              <Card key={i} card={e} idx={i} isTop={i === cards.length - 1} />
            ))}
          </div>
        </div>
      );
      
}