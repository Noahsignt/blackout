import { NavigateFunction } from 'react-router-dom';
import { isAuthenticated } from '../../../api/util';
import { createGame, joinGame } from '../../../api/game';
import type { CreateGameResponse } from '../../../types/game';

export interface GameForm {
  gameName: string;
}

export interface CreateGameHandlerProps {
  gameForm: GameForm;
  setGameForm: (form: GameForm) => void;
  setIsLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  setSuccess: (success: string | null) => void;
  navigate: NavigateFunction;
}

export const handleCreateGame = async ({
  gameForm,
  setGameForm,
  setIsLoading,
  setError,
  setSuccess,
  navigate
}: CreateGameHandlerProps) => {
  if (!isAuthenticated()) {
    setError('Please log in to create a game');
    return;
  }

  if (!gameForm.gameName.trim()) {
    setError('Please enter a game name');
    return;
  }

  setIsLoading(true);
  setError(null);
  setSuccess(null);

  try {
    const response: CreateGameResponse = await createGame({ numRounds: 10 });
    setGameForm({ gameName: '' });

    await joinGame(response.id);

    // redirect to the game page
    navigate(`/game/${response.id}`);
  } catch (err) {
    setError(err instanceof Error ? err.message : 'Failed to create game');
  } finally {
    setIsLoading(false);
  }
};