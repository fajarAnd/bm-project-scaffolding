import { useState, useEffect, FormEvent } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { Event } from '../types/event.types';
import { eventService } from '../services/event.service';
import { ticketService } from '../services/ticket.service';

export const PurchasePage = () => {
  const { user } = useAuth();
  const [events, setEvents] = useState<Event[]>([]);
  const [eventId, setEventId] = useState('');
  const [quantity, setQuantity] = useState(1);
  const [isLoading, setIsLoading] = useState(false);
  const [isLoadingEvents, setIsLoadingEvents] = useState(true);
  const [message, setMessage] = useState<{ type: 'success' | 'error'; text: string } | null>(null);

  useEffect(() => {
    loadEvents();
  }, []);

  const loadEvents = async () => {
    try {
      setIsLoadingEvents(true);
      const data = await eventService.getEvents();
      setEvents(data);
    } catch (err) {
      setMessage({
        type: 'error',
        text: 'Failed to load events. Please refresh the page.'
      });
    } finally {
      setIsLoadingEvents(false);
    }
  };

  const selectedEvent = events.find(e => e.id === eventId);
  const totalPrice = selectedEvent ? selectedEvent.ticket_price * quantity : 0;

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setMessage(null);

    if (!eventId) {
      setMessage({ type: 'error', text: 'Please select an event' });
      return;
    }

    if (quantity < 1) {
      setMessage({ type: 'error', text: 'Quantity must be at least 1' });
      return;
    }

    setIsLoading(true);

    try {
      const response = await ticketService.purchaseTicket({
        event_id: eventId,
        quantity: quantity
      });

      setMessage({
        type: 'success',
        text: response.message || `Successfully purchased ${quantity} ticket(s) for ${selectedEvent?.title}! Order ID: ${response.order_id}`
      });

      // Reset form
      setEventId('');
      setQuantity(1);

      // Reload events to get updated ticket availability
      loadEvents();
    } catch (err: any) {
      setMessage({
        type: 'error',
        text: err.response?.data?.error || err.message || 'Purchase failed. Please try again.'
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '600px', margin: '0 auto' }}>
      <h1 style={{ marginBottom: '10px' }}>Purchase Tickets</h1>
      <p style={{ color: '#666', marginBottom: '30px' }}>
        Welcome, {user?.email}! Select an event and purchase your tickets.
      </p>

      <form onSubmit={handleSubmit} style={{
        backgroundColor: 'white',
        padding: '30px',
        borderRadius: '8px',
        boxShadow: '0 2px 8px rgba(0,0,0,0.1)'
      }}>
        <div style={{ marginBottom: '20px' }}>
          <label htmlFor="event" style={{ display: 'block', marginBottom: '8px', fontWeight: '500' }}>
            Select Event
          </label>
          <select
            id="event"
            value={eventId}
            onChange={(e) => setEventId(e.target.value)}
            disabled={isLoading || isLoadingEvents}
            style={{
              width: '100%',
              padding: '10px',
              fontSize: '16px',
              border: '1px solid #ddd',
              borderRadius: '4px',
              backgroundColor: 'white'
            }}
          >
            <option value="">
              {isLoadingEvents ? 'Loading events...' : '-- Choose an event --'}
            </option>
            {events.map((event) => (
              <option key={event.id} value={event.id}>
                {event.title} (Rp{event.ticket_price.toLocaleString('id-ID')})
              </option>
            ))}
          </select>
        </div>

        <div style={{ marginBottom: '20px' }}>
          <label htmlFor="quantity" style={{ display: 'block', marginBottom: '8px', fontWeight: '500' }}>
            Quantity
          </label>
          <input
            id="quantity"
            type="number"
            min="1"
            max="10"
            value={quantity}
            onChange={(e) => setQuantity(parseInt(e.target.value) || 1)}
            disabled={isLoading}
            style={{
              width: '100%',
              padding: '10px',
              fontSize: '16px',
              border: '1px solid #ddd',
              borderRadius: '4px'
            }}
          />
        </div>

        {selectedEvent && (
          <div style={{
            padding: '15px',
            backgroundColor: '#f8f9fa',
            borderRadius: '4px',
            marginBottom: '20px'
          }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '5px' }}>
              <span>Price per ticket:</span>
              <span style={{ fontWeight: '500' }}>Rp{selectedEvent.ticket_price.toLocaleString('id-ID')}</span>
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '5px' }}>
              <span>Quantity:</span>
              <span style={{ fontWeight: '500' }}>{quantity}</span>
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '5px' }}>
              <span>Available tickets:</span>
              <span style={{ fontWeight: '500', color: selectedEvent.available_tickets < 10 ? '#dc3545' : '#666' }}>
                {selectedEvent.available_tickets}
              </span>
            </div>
            <hr style={{ margin: '10px 0', border: 'none', borderTop: '1px solid #ddd' }} />
            <div style={{ display: 'flex', justifyContent: 'space-between' }}>
              <span style={{ fontWeight: '600' }}>Total:</span>
              <span style={{ fontWeight: '600', fontSize: '18px', color: '#007bff' }}>
                Rp{totalPrice.toLocaleString('id-ID')}
              </span>
            </div>
          </div>
        )}

        {message && (
          <div style={{
            padding: '12px',
            marginBottom: '20px',
            borderRadius: '4px',
            backgroundColor: message.type === 'success' ? '#d4edda' : '#f8d7da',
            color: message.type === 'success' ? '#155724' : '#721c24',
            border: `1px solid ${message.type === 'success' ? '#c3e6cb' : '#f5c6cb'}`
          }}>
            {message.text}
          </div>
        )}

        <button
          type="submit"
          disabled={isLoading || isLoadingEvents || !eventId}
          style={{
            width: '100%',
            padding: '12px',
            fontSize: '16px',
            fontWeight: '500',
            backgroundColor: isLoading || isLoadingEvents || !eventId ? '#ccc' : '#28a745',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: isLoading || isLoadingEvents || !eventId ? 'not-allowed' : 'pointer',
            transition: 'background-color 0.2s'
          }}
        >
          {isLoading ? 'Processing...' : 'Purchase Tickets'}
        </button>
      </form>
    </div>
  );
};