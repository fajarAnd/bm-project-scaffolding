package models

import "time"

// RolePermission represents the many-to-many relationship between roles and permissions
// Allows dynamic permission assignment to roles without code changes
type RolePermission struct {
	ID           string    `db:"id" json:"id"`
	RoleID       string    `db:"role_id" json:"role_id"`
	PermissionID string    `db:"permission_id" json:"permission_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`

	// Relationships (for eager loading)
	Role       *Role       `db:"-" json:"role,omitempty"`
	Permission *Permission `db:"-" json:"permission,omitempty"`
}