import { Client } from 'boardgame.io/react';
import { Blackout } from './Game';

const App = Client({ game: Blackout });

export default App;