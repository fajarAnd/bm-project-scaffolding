import { useState, useEffect } from 'react';
import { EventCard } from '../components/EventCard';
import { Event } from '../types/event.types';
import { eventService } from '../services/event.service';

export const EventsPage = () => {
  const [events, setEvents] = useState<Event[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    loadEvents();
  }, []);

  const loadEvents = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await eventService.getEvents();
      setEvents(data);
    } catch (err: any) {
      setError(err.response?.data?.error || err.message || 'Failed to load events');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '1200px', margin: '0 auto' }}>
      <h1 style={{ marginBottom: '10px' }}>Available Events</h1>
      <p style={{ color: '#666', marginBottom: '30px' }}>
        Browse our upcoming events and book your tickets today!
      </p>

      {isLoading && (
        <div style={{ textAlign: 'center', padding: '40px', color: '#666' }}>
          <p>Loading events...</p>
        </div>
      )}

      {error && (
        <div style={{
          padding: '20px',
          backgroundColor: '#f8d7da',
          color: '#721c24',
          borderRadius: '4px',
          border: '1px solid #f5c6cb',
          marginBottom: '20px'
        }}>
          <strong>Error:</strong> {error}
          <button
            onClick={loadEvents}
            style={{
              marginLeft: '15px',
              padding: '5px 15px',
              backgroundColor: '#721c24',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            Retry
          </button>
        </div>
      )}

      {!isLoading && !error && events.length === 0 && (
        <div style={{
          textAlign: 'center',
          padding: '40px',
          backgroundColor: '#f8f9fa',
          borderRadius: '8px',
          color: '#666'
        }}>
          <p>No events available at the moment. Check back later!</p>
        </div>
      )}

      {!isLoading && !error && events.length > 0 && (
        <div>
          {events.map((event) => (
            <EventCard key={event.id} event={event} />
          ))}
        </div>
      )}
    </div>
  );
};