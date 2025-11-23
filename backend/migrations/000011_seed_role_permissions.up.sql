-- Seed role-permission mappings

-- Admin: all permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'admin';

-- User: basic permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'user'
AND p.name IN ('events.read', 'tickets.purchase', 'tickets.read');

-- Organizer: event management
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'organizer'
AND p.name IN ('events.create', 'events.read', 'events.update', 'tickets.read');

-- Validator: ticket validation only
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'validator'
AND p.name IN ('tickets.validate', 'tickets.read');