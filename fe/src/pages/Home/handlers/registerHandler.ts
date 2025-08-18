import { signUp } from '../../../api/auth';
import type { SignUpResponse } from '../../../types/auth';

export interface RegisterForm {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export interface RegisterHandlerProps {
  registerForm: RegisterForm;
  setRegisterForm: (form: RegisterForm) => void;
  setIsLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  setSuccess: (success: string | null) => void;
  setActiveTab: (tab: string) => void;
}

export const handleRegister = async ({
  registerForm,
  setRegisterForm,
  setIsLoading,
  setError,
  setSuccess,
  setActiveTab
}: RegisterHandlerProps) => {
  if (!registerForm.username || !registerForm.password || !registerForm.confirmPassword) {
    setError('Please fill in all fields');
    return;
  }

  if (registerForm.password !== registerForm.confirmPassword) {
    setError('Passwords do not match');
    return;
  }

  if (registerForm.password.length < 6) {
    setError('Password must be at least 6 characters');
    return;
  }

  setIsLoading(true);
  setError(null);
  setSuccess(null);

  try {
    const response: SignUpResponse = await signUp(registerForm.username, registerForm.password);
    setSuccess('Registration successful! You can now log in.');
    setRegisterForm({ username: '', email: '', password: '', confirmPassword: '' });
    setActiveTab('login');
  } catch (err) {
    setError(err instanceof Error ? err.message : 'Registration failed');
  } finally {
    setIsLoading(false);
  }
};