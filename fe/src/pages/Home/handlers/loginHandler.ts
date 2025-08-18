import { login } from '../../../api/auth';
import type { LoginResponse } from '../../../types/auth';

export interface LoginForm {
  username: string;
  password: string;
}

export interface LoginHandlerProps {
  loginForm: LoginForm;
  setLoginForm: (form: LoginForm) => void;
  setIsLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  setSuccess: (success: string | null) => void;
}

export const handleLogin = async ({
  loginForm,
  setLoginForm,
  setIsLoading,
  setError,
  setSuccess
}: LoginHandlerProps) => {
  if (!loginForm.username || !loginForm.password) {
    setError('Please fill in all fields');
    return;
  }

  setIsLoading(true);
  setError(null);
  setSuccess(null);

  try {
    const response: LoginResponse = await login(loginForm.username, loginForm.password);
    setSuccess('Login successful!');
    setLoginForm({ username: '', password: '' });
  } catch (err) {
    setError(err instanceof Error ? err.message : 'Login failed');
  } finally {
    setIsLoading(false);
  }
};