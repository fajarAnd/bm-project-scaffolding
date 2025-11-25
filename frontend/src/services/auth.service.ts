import { apiClient } from '../utils/api';
import { storage } from '../utils/storage';
import { AuthResponse, User } from '../types/auth.types';

export const authService = {
  async login(email: string, password: string): Promise<AuthResponse> {
    const response = await apiClient.post<{ success: boolean; data: AuthResponse }>('/auth/login', {
      email,
      password,
    });

    // Backend returns nested response: { success, data: { token, user } }
    const authData = response.data.data;

    // Save token to storage
    storage.setToken(authData.token);

    return authData;
  },

  logout(): void {
    storage.removeToken();
  },

  async getCurrentUser(): Promise<User | null> {
    if (!storage.hasToken()) {
      return null;
    }

    try {
      const response = await apiClient.get<{ success: boolean; data: User }>('/users/me');
      return response.data.data || response.data;
    } catch (error) {
      storage.removeToken();
      return null;
    }
  },

  isAuthenticated(): boolean {
    return storage.hasToken();
  }
};