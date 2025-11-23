package models

import "time"

// Event represents a live event (concert, conference, etc.)
type Event struct {
	ID               string    `db:"id" json:"id"`
	Title            string    `db:"title" json:"title"`
	Description      string    `db:"description" json:"description"`
	EventDate        time.Time `db:"event_date" json:"event_date"`
	Venue            string    `db:"venue" json:"venue"`
	TicketPrice      float64   `db:"ticket_price" json:"ticket_price"`
	TotalTickets     int       `db:"total_tickets" json:"total_tickets"`
	AvailableTickets int       `db:"available_tickets" json:"available_tickets"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
