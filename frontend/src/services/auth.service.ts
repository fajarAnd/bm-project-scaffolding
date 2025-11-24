import { apiClient } from '../utils/api';
import { storage } from '../utils/storage';
import { AuthResponse, LoginRequest, User } from '../types/auth.types';

export const authService = {
  async login(email: string, password: string): Promise<AuthResponse> {
    const response = await apiClient.post<AuthResponse>('/auth/login', {
      email,
      password,
    });

    // Save token to storage
    storage.setToken(response.data.token);

    return response.data;
  },

  logout(): void {
    storage.removeToken();
  },

  async getCurrentUser(): Promise<User | null> {
    if (!storage.hasToken()) {
      return null;
    }

    try {
      const response = await apiClient.get<User>('/users/me');
      return response.data;
    } catch (error) {
      // Token might be invalid
      storage.removeToken();
      return null;
    }
  },

  isAuthenticated(): boolean {
    return storage.hasToken();
  }
};