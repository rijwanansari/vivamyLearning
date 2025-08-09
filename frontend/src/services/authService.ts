import api from './api';
import { AuthResponse, User } from '../types';
import Cookies from 'js-cookie';

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface RegisterData {
  name: string;
  email: string;
  password: string;
}

class AuthService {
  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>('/auth/login', credentials);
    const { data } = response.data;
    
    // Store tokens and user info
    Cookies.set('access_token', data.access_token, { expires: 7 });
    Cookies.set('user', JSON.stringify(data.user), { expires: 7 });
    
    return response.data;
  }

  async register(userData: RegisterData): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>('/auth/register', userData);
    const { data } = response.data;
    
    // Store tokens and user info
    Cookies.set('access_token', data.access_token, { expires: 7 });
    Cookies.set('user', JSON.stringify(data.user), { expires: 7 });
    
    return response.data;
  }

  logout(): void {
    Cookies.remove('access_token');
    Cookies.remove('user');
    window.location.href = '/login';
  }

  getCurrentUser(): User | null {
    const userCookie = Cookies.get('user');
    return userCookie ? JSON.parse(userCookie) : null;
  }

  isAuthenticated(): boolean {
    return !!Cookies.get('access_token');
  }

  getToken(): string | undefined {
    return Cookies.get('access_token');
  }
}

export default new AuthService();
