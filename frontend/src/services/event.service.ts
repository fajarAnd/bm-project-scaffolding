import { apiClient } from '../utils/api';
import { Event, EventListResponse } from '../types/event.types';
import { APIResponse } from '../types/api.types';

export const eventService = {
  async getEvents(): Promise<Event[]> {
    const response = await apiClient.get<APIResponse<EventListResponse>>('/events');
    return response.data.data?.events || [];
  },
    
  async getEventById(id: string): Promise<Event> {
    const response = await apiClient.get<APIResponse<{ event: Event }>>(`/events/${id}`);
    if (!response.data.data?.event) {
      throw new Error('Event not found');
    }
    return response.data.data.event;
  },
};