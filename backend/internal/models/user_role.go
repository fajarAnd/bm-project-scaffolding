package models

import "time"

// UserRole represents the many-to-many relationship between users and roles
// Allows users to have multiple roles (e.g., user + organizer)
type UserRole struct {
	ID         string    `db:"id" json:"id"`
	UserID     string    `db:"user_id" json:"user_id"`
	RoleID     string    `db:"role_id" json:"role_id"`
	AssignedAt time.Time `db:"assigned_at" json:"assigned_at"`
	AssignedBy *string   `db:"assigned_by" json:"assigned_by,omitempty"` // nullable, FK to users.id

	// Relationships (for eager loading)
	User *User `db:"-" json:"user,omitempty"`
	Role *Role `db:"-" json:"role,omitempty"`
}