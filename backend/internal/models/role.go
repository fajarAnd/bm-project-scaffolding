package models

import "time"

// Role represents a system role (admin, user, organizer, etc.)
type Role struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	// Relationships (loaded via joins)
	Permissions []*Permission `db:"-" json:"permissions,omitempty"`
}

// Common role names
const (
	RoleAdmin     = "admin"
	RoleUser      = "user"
	RoleOrganizer = "organizer"
	RoleValidator = "validator"
)