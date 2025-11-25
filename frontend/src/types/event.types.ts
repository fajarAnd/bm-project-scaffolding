export interface Event {
  id: string;
  title: string;
  description: string;
  event_date: string;
  venue: string;
  ticket_price: number;
  total_tickets: number;
  available_tickets: number;
  created_at: string;
  updated_at: string;
}

export interface EventListResponse {
  events: Event[];
  total: number;
  page: number;
}