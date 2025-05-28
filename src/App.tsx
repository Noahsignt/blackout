import { Client } from 'boardgame.io/react';
import { Blackout } from './game/Game';
import { Board } from './components/board/Board';

const App = Client({ 
  game: Blackout,
  board: Board
});

export default App;