export const API_URL = process.env.NODE_ENV === 'production' 
  ? 'http://blackout-be.duckdns.org:8080'
  : 'http://localhost:8080';


export const apiRequest = async (endpoint: string, options: RequestInit = {}) => {
  const url = `${API_URL}${endpoint}`;
  
  const defaultHeaders = {
    'Content-Type': 'application/json',
  };

  const config: RequestInit = {
    ...options,
    headers: {
      ...defaultHeaders,
      ...options.headers,
    },
  };

  const response = await fetch(url, config);
  
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || `HTTP error! status: ${response.status}`);
  }

  if (response.status === 204 || response.headers.get('content-length') === '0') {
    return null;
  }

  return response.json();
};

export const authenticatedRequest = async (endpoint: string, options: RequestInit = {}) => {
  const token = getAuthToken();
  
  if (!token) {
    throw new Error('No authentication token found');
  }

  return apiRequest(endpoint, {
    ...options,
    headers: {
      ...options.headers,
      'Authorization': `Bearer ${token}`,
    },
  });
};

// token management utils
export const getAuthToken = (): string | null => {
  return localStorage.getItem('blackout_auth_token');
};

export const setAuthToken = (token: string): void => {
  localStorage.setItem('blackout_auth_token', token);
};

export const removeAuthToken = (): void => {
  localStorage.removeItem('blackout_auth_token');
};

export const isAuthenticated = (): boolean => {
  return !!getAuthToken();
};