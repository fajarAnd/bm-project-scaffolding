import { apiClient } from '../utils/api';
import { PurchaseRequest, PurchaseResponse } from '../types/ticket.types';
import { APIResponse } from '../types/api.types';

export const ticketService = {
  async purchaseTicket(request: PurchaseRequest): Promise<PurchaseResponse> {
    const response = await apiClient.post<APIResponse<PurchaseResponse>>('/tickets/purchase', request);
    if (!response.data.data) {
      throw new Error(response.data.error || 'Purchase failed');
    }
    return response.data.data;
  },
};