import { writable } from 'svelte/store';

interface User {
  id: string;
  username: string;
  email: string;
}

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
}

function createAuthStore() {
  const initialState: AuthState = {
    user: null,
    token: null,
    isAuthenticated: false
  };

  const { subscribe, set, update } = writable<AuthState>(initialState);

  return {
    subscribe,
    login: async (username: string, password: string) => {
      try {
        // In production, this would be a real API call
        // Mock implementation for demo
        if (username === 'admin' && password === 'password') {
          const mockUser = { id: '1', username, email: 'admin@example.com' };
          const mockToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIn0.mock-token';
          
          localStorage.setItem('musicAppToken', mockToken);
          localStorage.setItem('musicAppUser', JSON.stringify(mockUser));
          
          set({ user: mockUser, token: mockToken, isAuthenticated: true });
          return { success: true };
        }
        return { success: false, message: 'Invalid credentials' };
      } catch (error) {
        console.error('Login error:', error);
        return { success: false, message: 'Authentication failed' };
      }
    },
    logout: () => {
      localStorage.removeItem('musicAppToken');
      localStorage.removeItem('musicAppUser');
      set(initialState);
    },
    checkAuth: async () => {
      console.log('sexo tilin');
      const token = localStorage.getItem('musicAppToken');
      const userStr = localStorage.getItem('musicAppUser');
      
      if (token && userStr) {
        try {
          const user = JSON.parse(userStr);
          set({ user, token, isAuthenticated: true });
        } catch (e) {
          localStorage.removeItem('musicAppToken');
          localStorage.removeItem('musicAppUser');
        }
      }
    }
  };
}

export const authStore = createAuthStore();
