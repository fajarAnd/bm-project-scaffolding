package models

// TicketOrder represents a ticket purchase order
type TicketOrder struct {
	// TODO: Define ticket order fields
	// Expected fields:
	// - ID (UUID)
	// - EventID (FK to events)
	// - UserID (FK to users)
	// - Quantity
	// - TotalPrice
	// - Status (pending, paid, confirmed, cancelled, refunded)
	// - PaymentID
	// - CreatedAt
	// - UpdatedAt
}

// Ticket represents an individual ticket
type Ticket struct {
	// TODO: Define ticket fields
	// Expected fields:
	// - ID (UUID)
	// - OrderID (FK to ticket_orders)
	// - TicketCode (unique)
	// - Status (valid, used, cancelled, transferred)
	// - UsedAt
	// - CreatedAt
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
