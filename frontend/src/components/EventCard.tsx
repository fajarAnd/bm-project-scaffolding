import { Event } from '../types/event.types';

interface EventCardProps {
  event: Event;
  onSelectEvent?: (eventId: string) => void;
}

export const EventCard: React.FC<EventCardProps> = ({ event, onSelectEvent }) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      weekday: 'short',
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  };

  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(price);
  };

  return (
    <div style={{
      border: '1px solid #ddd',
      borderRadius: '8px',
      padding: '20px',
      marginBottom: '15px',
      backgroundColor: 'white',
      boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
      transition: 'transform 0.2s, box-shadow 0.2s',
      cursor: onSelectEvent ? 'pointer' : 'default'
    }}
    onClick={() => onSelectEvent?.(event.id)}
    onMouseEnter={(e) => {
      if (onSelectEvent) {
        e.currentTarget.style.transform = 'translateY(-2px)';
        e.currentTarget.style.boxShadow = '0 4px 8px rgba(0,0,0,0.15)';
      }
    }}
    onMouseLeave={(e) => {
      if (onSelectEvent) {
        e.currentTarget.style.transform = 'translateY(0)';
        e.currentTarget.style.boxShadow = '0 2px 4px rgba(0,0,0,0.1)';
      }
    }}
    >
      <h3 style={{ margin: '0 0 10px 0', color: '#333' }}>{event.title}</h3>
      <p style={{ margin: '0 0 10px 0', color: '#666', fontSize: '14px' }}>
        {event.description}
      </p>

      <div style={{ display: 'flex', gap: '20px', flexWrap: 'wrap', marginTop: '15px' }}>
        <div>
          <span style={{ color: '#666', fontSize: '13px' }}>📅 Date:</span>
          <p style={{ margin: '5px 0 0 0', fontWeight: '500' }}>
            {formatDate(event.event_date)}
          </p>
        </div>

        <div>
          <span style={{ color: '#666', fontSize: '13px' }}>📍 Location:</span>
          <p style={{ margin: '5px 0 0 0', fontWeight: '500' }}>
            {event.venue}
          </p>
        </div>

        <div>
          <span style={{ color: '#666', fontSize: '13px' }}>💵 Price:</span>
          <p style={{ margin: '5px 0 0 0', fontWeight: '500', color: '#007bff' }}>
            {formatPrice(event.ticket_price)}
          </p>
        </div>

        <div>
          <span style={{ color: '#666', fontSize: '13px' }}>🎟️ Available:</span>
          <p style={{ margin: '5px 0 0 0', fontWeight: '500' }}>
            {event.available_tickets} tickets
          </p>
        </div>
      </div>
    </div>
  );
};