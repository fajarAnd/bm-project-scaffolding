-- Events table
CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date TIMESTAMP NOT NULL,
    venue VARCHAR(255) NOT NULL,
    ticket_price DECIMAL(10,2) NOT NULL,
    total_tickets INTEGER NOT NULL,
    available_tickets INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT available_tickets_check CHECK (available_tickets >= 0),
    CONSTRAINT ticket_price_check CHECK (ticket_price >= 0)
);

-- Indexes
CREATE INDEX idx_events_event_date ON events(event_date);
CREATE INDEX idx_events_available_tickets ON events(available_tickets);