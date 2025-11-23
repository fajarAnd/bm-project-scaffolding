package models

// User represents a user in the system
type User struct {
	// TODO: Define user fields
	// Expected fields:
	// - ID (UUID)
	// - Email (unique)
	// - PasswordHash
	// - Role (user/admin)
	// - CreatedAt
	// - UpdatedAt
}

// UserRole defines user role types
type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)
