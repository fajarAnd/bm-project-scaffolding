package models

import "time"

// TicketOrder represents a ticket purchase order
type TicketOrder struct {
	ID         string    `db:"id" json:"id"`
	EventID    string    `db:"event_id" json:"event_id"`
	UserID     string    `db:"user_id" json:"user_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	TotalPrice float64   `db:"total_price" json:"total_price"`
	Status     string    `db:"status" json:"status"`
	PaymentID  *string   `db:"payment_id" json:"payment_id,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`

	// Relationships (loaded via joins)
	Event   *Event   `db:"-" json:"event,omitempty"`
	User    *User    `db:"-" json:"user,omitempty"`
	Tickets []Ticket `db:"-" json:"tickets,omitempty"`
}

// Ticket represents an individual ticket
type Ticket struct {
	ID         string     `db:"id" json:"id"`
	OrderID    string     `db:"order_id" json:"order_id"`
	TicketCode string     `db:"ticket_code" json:"ticket_code"`
	Status     string     `db:"status" json:"status"`
	UsedAt     *time.Time `db:"used_at" json:"used_at,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`

	// Relationships (loaded via joins)
	Order *TicketOrder `db:"-" json:"order,omitempty"`
}

// TicketOrderStatus defines ticket order status types
type TicketOrderStatus string

const (
	OrderStatusPending   TicketOrderStatus = "pending"
	OrderStatusPaid      TicketOrderStatus = "paid"
	OrderStatusConfirmed TicketOrderStatus = "confirmed"
	OrderStatusCancelled TicketOrderStatus = "cancelled"
	OrderStatusRefunded  TicketOrderStatus = "refunded"
)

// TicketStatus defines ticket status types
type TicketStatus string

const (
	TicketStatusValid       TicketStatus = "valid"
	TicketStatusUsed        TicketStatus = "used"
	TicketStatusCancelled   TicketStatus = "cancelled"
	TicketStatusTransferred TicketStatus = "transferred"
)
