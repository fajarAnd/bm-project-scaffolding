export interface PurchaseRequest {
  event_id: string;
  quantity: number;
}

export interface PurchaseResponse {
  order_id: string;
  transaction_id: string;
  status: string;
  message?: string;
}

export interface TicketOrder {
  id: string;
  event_id: string;
  user_id: string;
  quantity: number;
  total_price: number;
  status: string;
  payment_id?: string;
  created_at: string;
  updated_at: string;
}