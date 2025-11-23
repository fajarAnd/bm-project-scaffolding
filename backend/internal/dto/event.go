package dto

import "github.com/baramulti/ticketing-system/backend/internal/models"

type CreateEventRequest struct {
	Title            string  `json:"title" binding:"required"`
	Description      string  `json:"description"`
	EventDate        string  `json:"event_date" binding:"required"`
	Venue            string  `json:"venue" binding:"required"`
	TicketPrice      float64 `json:"ticket_price" binding:"required,min=0"`
	TotalTickets     int     `json:"total_tickets" binding:"required,min=1"`
	AvailableTickets int     `json:"available_tickets" binding:"required,min=0"`
}

type UpdateEventRequest struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	EventDate   string  `json:"event_date,omitempty"`
	Venue       string  `json:"venue,omitempty"`
	TicketPrice float64 `json:"ticket_price,omitempty"`
}

type EventResponse struct {
	Event *models.Event `json:"event"`
}

type EventListResponse struct {
	Events []*models.Event `json:"events"`
	Total  int             `json:"total"`
	Page   int             `json:"page"`
}
