import { EventCard } from '../components/EventCard';
import { Event } from '../types/event.types';

// Mock events data - will be replaced with API call in Phase 4
const mockEvents: Event[] = [
  {
    id: '1',
    name: 'Summer Music Festival 2025',
    description: 'Join us for the biggest music festival of the year featuring top artists from around the world.',
    date: '2025-07-15T18:00:00Z',
    location: 'Ice BSD, Tangerang',
    price: 150.00,
    available_tickets: 500,
    created_at: '2025-01-01T00:00:00Z',
    updated_at: '2025-01-01T00:00:00Z'
  },
  {
    id: '2',
    name: 'Tech Conference 2025',
    description: 'Discover the latest trends in technology and network with industry leaders.',
    date: '2025-08-20T09:00:00Z',
    location: 'Ice BSD, Tangerang',
    price: 299.00,
    available_tickets: 200,
    created_at: '2025-01-02T00:00:00Z',
    updated_at: '2025-01-02T00:00:00Z'
  },
  {
    id: '3',
    name: 'Food & Wine Festival',
    description: 'Experience culinary delights from world-renowned chefs and sommeliers.',
    date: '2025-09-10T12:00:00Z',
    location: 'Ice BSD, Tangerang',
    price: 85.00,
    available_tickets: 150,
    created_at: '2025-01-03T00:00:00Z',
    updated_at: '2025-01-03T00:00:00Z'
  }
];

export const EventsPage = () => {
  return (
    <div style={{ padding: '20px', maxWidth: '1200px', margin: '0 auto' }}>
      <h1 style={{ marginBottom: '10px' }}>Available Events</h1>
      <p style={{ color: '#666', marginBottom: '30px' }}>
        Browse our upcoming events and book your tickets today!
      </p>

      <div>
        {mockEvents.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>

      <p style={{ marginTop: '30px', fontSize: '14px', color: '#999', fontStyle: 'italic' }}>
        Note: Currently showing mock data. Real-time event data will be loaded from API in Phase 4.
      </p>
    </div>
  );
};