import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { startGame } from '../../api/game';
import type { GameResponse } from '../../types/game';

export default function GamePage() {
  const { id } = useParams<{ id: string }>();
  const [game, setGame] = useState<GameResponse | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

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
      {/* Background styling similar to home page */}
      <div className="absolute inset-0 bg-gradient-to-br from-stone-600 via-stone-700 to-stone-800"></div>
      
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