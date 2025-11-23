-- Seed default roles
INSERT INTO roles (name, description) VALUES
    ('admin', 'Full system access'),
    ('user', 'Basic customer role'),
    ('organizer', 'Event organizer role'),
    ('validator', 'Ticket validator role');