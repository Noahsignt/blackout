import { NavigateFunction } from 'react-router-dom';
import { isAuthenticated } from '../../../api/util';
import { joinGame } from '../../../api/game';
import type { GameResponse } from '../../../types/game';

export interface JoinGameForm {
  gameId: string;
}

export interface JoinGameHandlerProps {
  joinGameForm: JoinGameForm;
  setJoinGameForm: (form: JoinGameForm) => void;
  setIsLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  setSuccess: (success: string | null) => void;
  navigate: NavigateFunction;
}

export const handleJoinGame = async ({
  joinGameForm,
  setJoinGameForm,
  setIsLoading,
  setError,
  setSuccess,
  navigate
}: JoinGameHandlerProps) => {
  if (!isAuthenticated()) {
    setError('Please log in to join a game');
    return;
  }

  if (!joinGameForm.gameId.trim()) {
    setError('Please enter a game ID');
    return;
  }

  setIsLoading(true);
  setError(null);
  setSuccess(null);

  try {
    const response: GameResponse = await joinGame(joinGameForm.gameId.trim());
    setJoinGameForm({ gameId: '' });
    // redirect to the game page
    navigate(`/game/${response.id}`);
  } catch (err) {
    setError(err instanceof Error ? err.message : 'Failed to join game');
  } finally {
    setIsLoading(false);
  }
};