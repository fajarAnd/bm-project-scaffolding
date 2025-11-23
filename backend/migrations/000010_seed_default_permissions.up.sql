-- Seed default permissions
INSERT INTO permissions (name, resource, action, description) VALUES
    -- Event permissions
    ('events.create', 'events', 'create', 'Create new events'),
    ('events.read', 'events', 'read', 'View events'),
    ('events.update', 'events', 'update', 'Update events'),
    ('events.delete', 'events', 'delete', 'Delete events'),

    -- User permissions
    ('users.create', 'users', 'create', 'Create new users'),
    ('users.read', 'users', 'read', 'View users'),
    ('users.update', 'users', 'update', 'Update users'),
    ('users.delete', 'users', 'delete', 'Delete users'),

    -- Ticket permissions
    ('tickets.purchase', 'tickets', 'purchase', 'Purchase tickets'),
    ('tickets.read', 'tickets', 'read', 'View tickets'),
    ('tickets.validate', 'tickets', 'validate', 'Validate tickets at venue'),
    ('tickets.refund', 'tickets', 'refund', 'Refund tickets');