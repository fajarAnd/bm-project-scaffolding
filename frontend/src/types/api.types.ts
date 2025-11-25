// Generic API response wrapper from backend
export interface APIResponse<T> {
  success: boolean;
  data?: T;
  error?: string;
}