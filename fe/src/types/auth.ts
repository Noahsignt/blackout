// Auth request/response types matching backend DTOs
export interface AuthRequest {
  username: string;
  password: string;
}

export interface SignUpResponse {
  id: string;
  username: string;
  createdAt: string;
}

export interface LoginResponse {
  authString: string;
}

export interface AuthError {
  message: string;
  status?: number;
}

export interface ChangePasswordRequest {
  newPassword: string;
}

export interface UpdateImageRequest {
  imageUrl: string;
}

// User type matching backend model
export interface User {
  id: string;
  username: string;
  image_url?: string;
  createdAt: string;
}

// Auth state types
export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  token: string | null;
}