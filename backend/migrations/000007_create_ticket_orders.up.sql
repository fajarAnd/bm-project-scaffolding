-- Ticket orders table
CREATE TABLE ticket_orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE RESTRICT,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    payment_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT quantity_check CHECK (quantity > 0)
);

-- Indexes
CREATE INDEX idx_ticket_orders_user_id ON ticket_orders(user_id);
CREATE INDEX idx_ticket_orders_event_id ON ticket_orders(event_id);
CREATE INDEX idx_ticket_orders_status ON ticket_orders(status);
CREATE INDEX idx_ticket_orders_created_at ON ticket_orders(created_at DESC);