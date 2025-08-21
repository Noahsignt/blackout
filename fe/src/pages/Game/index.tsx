import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { startGame, getGame } from '../../api/game';
import { useNavigate } from 'react-router-dom';
import type { GameResponse, Player } from '../../types/game';

export default function GamePage() {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [game, setGame] = useState<GameResponse | null>(null);
  const [players, setPlayers] = useState<Array<Player>>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if(!id) {
      navigate('/');
      return;
    }
  
    const fetchGame = async () => {
      try {
        const gameData = await getGame(id!);
        // setGame(gameData);
        setPlayers(gameData.players);
      } catch (err) {
        console.error(err);
        setError('Failed to fetch game data');
      }
    };

    fetchGame();
    
    const pollInterval = setInterval(fetchGame, 5000);
    return () => clearInterval(pollInterval);
  }, [id, navigate])

  const handleStartGame = async () => {
    if (!id) return;
    
    setIsLoading(true);
    setError(null);

    try {
      const response = await startGame(id);
      setGame(response);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to start game');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-stone-700 relative">
      {/* Dingy concrete wall texture */}
      <div className="absolute inset-0 bg-gradient-to-br from-stone-600 via-stone-700 to-stone-800"></div>
      
      {/* Water stains and discoloration */}
      <div className="absolute inset-0" style={{
        backgroundImage: `
          radial-gradient(ellipse at 20% 30%, rgba(120, 113, 108, 0.4) 0%, transparent 60%),
          radial-gradient(ellipse at 80% 10%, rgba(87, 83, 81, 0.3) 0%, transparent 50%),
          radial-gradient(ellipse at 60% 80%, rgba(168, 162, 158, 0.2) 0%, transparent 40%)
        `,
        backgroundSize: '400px 300px, 300px 200px, 500px 400px'
      }}></div>
      
      {/* Flickering fluorescent light effect */}
      <div className="absolute top-0 left-1/4 right-1/4 h-2 bg-gradient-to-r from-transparent via-stone-300 to-transparent opacity-30"></div>
      
      {/* Cigarette smoke trails */}
      <div className="absolute top-1/3 left-1/4 w-px h-16 bg-gradient-to-t from-stone-400 to-transparent opacity-40 blur-sm"></div>
      <div className="absolute top-2/3 right-1/3 w-px h-12 bg-gradient-to-t from-stone-500 to-transparent opacity-30 blur-sm"></div>
      <div className="absolute bottom-1/4 left-2/3 w-px h-8 bg-gradient-to-t from-stone-400 to-transparent opacity-35 blur-sm"></div>
      
      <div className="relative z-10 container mx-auto px-6 py-12">
        <div className="text-center mb-12">
          <div className="inline-block bg-stone-800 px-8 py-4 border-2 border-stone-600 shadow-lg">
            <h1 className="text-4xl font-black text-stone-200 tracking-wider" 
                style={{ fontFamily: 'monospace' }}>
              GAME {id}
            </h1>
          </div>
        </div>

        {/* Error Message */}
        {error && (
          <div className="max-w-sm mx-auto mb-6">
            <div className="p-3 border-2 bg-red-900/50 border-red-600 text-red-200 font-mono text-sm">
              {error}
            </div>
          </div>
        )}

        {/* Game Info */}
        {game && (
          <div className="max-w-md mx-auto mb-8">
            <div className="bg-stone-800 border-2 border-stone-600 shadow-2xl p-6">
              <h2 className="text-lg font-black text-stone-200 mb-4 uppercase tracking-wide font-mono">
                Game Status
              </h2>
              <div className="space-y-2 text-stone-300 font-mono text-sm">
                <div>Rounds: {game.numRounds}</div>
                <div>Current Round: {game.round.number}</div>
                <div>Players: {game.players.length}</div>
                <div>Current Turn: Player {game.round.hand.turn + 1}</div>
              </div>
            </div>
          </div>
        )}

        {/* Players List */}
        {!game && (
          <div className="max-w-sm mx-auto mb-8">
            <div className="bg-stone-800 border-2 border-stone-600 shadow-2xl p-6">
              <h2 className="text-lg font-black text-stone-200 mb-4 uppercase tracking-wide font-mono">
                Players ({players.length})
              </h2>
              {players.length > 0 ? (
                <div className="space-y-2">
                  {players.map((player, index) => (
                    <div
                      key={player.id}
                      className="bg-stone-700 border border-stone-600 p-3 text-stone-300 font-mono text-sm"
                    >
                      <div className="flex justify-between items-center">
                        <span>{player.name}</span>
                        <span className="text-stone-500">#{index + 1}</span>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="text-stone-400 font-mono text-sm text-center py-4">
                  No players yet...
                </div>
              )}
            </div>
          </div>
        )}

        {/* Start Game Button */}
        {!game && (
          <div className="max-w-sm mx-auto">
            <div className="bg-stone-800 border-2 border-stone-600 shadow-2xl p-6">
              <h2 className="text-lg font-black text-stone-200 mb-6 uppercase tracking-wide font-mono">
                Ready to Start?
              </h2>
              <button
                onClick={handleStartGame}
                disabled={isLoading}
                className="w-full bg-stone-600 hover:bg-stone-500 text-stone-100 font-black py-3 uppercase tracking-wider text-sm transition-colors font-mono border border-stone-500 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isLoading ? 'STARTING...' : 'START GAME'}
              </button>
            </div>
          </div>
        )}

        {/* Game Board Placeholder */}
        {game && (
          <div className="max-w-4xl mx-auto">
            <div className="bg-stone-800 border-2 border-stone-600 shadow-2xl p-8">
              <h2 className="text-lg font-black text-stone-200 mb-6 uppercase tracking-wide font-mono text-center">
                Game Board
              </h2>
              <div className="text-center text-stone-400 font-mono">
                Game board implementation coming soon...
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}