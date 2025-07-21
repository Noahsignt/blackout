import { apiRequest, setAuthToken, removeAuthToken, getAuthToken, isAuthenticated } from './util';
import type { SignUpResponse, LoginResponse } from '../types/auth';

export const signUp = async (username: string, password: string): Promise<SignUpResponse> => {
  try {
    const response = await apiRequest('/users/signup', {
      method: 'POST',
      body: JSON.stringify({ username, password }),
    });
    
    return response as SignUpResponse;
  } catch (error) {
    throw new Error(error instanceof Error ? error.message : 'Signup failed');
  }
};

export const login = async (username: string, password: string): Promise<LoginResponse> => {
  try {
    const response = await apiRequest('/users/login', {
      method: 'POST',
      body: JSON.stringify({ username, password }),
    });
    
    // store auth token in local storage
    if (response.authString) {
      setAuthToken(response.authString);
    }
    
    return response as LoginResponse;
  } catch (error) {
    throw new Error(error instanceof Error ? error.message : 'Login failed');
  }
};

export const logout = (): void => {
  removeAuthToken();
};

export const validateToken = (): boolean => {
  const token = getAuthToken();
  
  if (!token) {
    return false;
  }
  
  try {
    // check JWT conforms to schema (header.payload.signature)
    const parts = token.split('.');
    if (parts.length !== 3) {
      removeAuthToken(); // remove invalid token
      return false;
    }
    
    // check expiration
    const payload = JSON.parse(atob(parts[1]));
    if (payload.exp && payload.exp * 1000 < Date.now()) {
      removeAuthToken(); // remove expired token
      return false;
    }
    
    return true;
  } catch (error) {
    removeAuthToken(); // remove malformed token
    return false;
  }
};

// Function to get user info from token (if needed)
export const getUserFromToken = (): any | null => {
  const token = getAuthToken();
  
  if (!token || !validateToken()) {
    return null;
  }
  
  try {
    const parts = token.split('.');
    const payload = JSON.parse(atob(parts[1]));
    return payload;
  } catch (error) {
    return null;
  }
};

// Auto-logout if token is invalid
export const ensureValidAuth = (): boolean => {
  if (!validateToken()) {
    logout();
    return false;
  }
  return true;
};