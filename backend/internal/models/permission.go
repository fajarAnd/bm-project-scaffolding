package models

import "time"

type Permission struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Resource    string    `db:"resource" json:"resource"`
	Action      string    `db:"action" json:"action"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// Common permission names (resource.action format)
const (
	// Event permissions
	PermEventCreate = "events.create"
	PermEventRead   = "events.read"
	PermEventUpdate = "events.update"
	PermEventDelete = "events.delete"

	// User permissions
	PermUserCreate = "users.create"
	PermUserRead   = "users.read"
	PermUserUpdate = "users.update"
	PermUserDelete = "users.delete"

	// Ticket permissions
	PermTicketPurchase = "tickets.purchase"
	PermTicketRead     = "tickets.read"
	PermTicketValidate = "tickets.validate"
	PermTicketRefund   = "tickets.refund"
)