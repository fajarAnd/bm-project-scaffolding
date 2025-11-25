package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

// EventRepository defines data access methods for events
type EventRepository interface {
	FindByID(ctx context.Context, id string) (*models.Event, error)
	List(ctx context.Context, limit, offset int) ([]*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
	Update(ctx context.Context, event *models.Event) error
	Delete(ctx context.Context, id string) error
	DecrementAvailableTickets(ctx context.Context, eventID string, quantity int) error
}

type eventRepository struct {
	db *sqlx.DB
}

// NewEventRepository creates a new event repository instance
func NewEventRepository(db *sqlx.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FindByID(ctx context.Context, id string) (*models.Event, error) {
	// TODO: implement actual database query
	// Example: SELECT * FROM events WHERE id = $1
	return nil, fmt.Errorf("not implemented")
}

func (r *eventRepository) List(ctx context.Context, limit, offset int) ([]*models.Event, error) {
	// TODO: implement event listing with pagination
	// Example: SELECT * FROM events ORDER BY event_date ASC LIMIT $1 OFFSET $2

	// mocking data from database
	var MockEvents = []*models.Event{
		{
			ID:               "evt-001",
			Title:            "Jakarta Tech Conference 2025",
			Description:      "A gathering of software engineers, architects, and tech leaders in Indonesia.",
			EventDate:        time.Date(2025, 3, 14, 9, 0, 0, 0, time.Local),
			Venue:            "Jakarta Convention Center",
			TicketPrice:      250000,
			TotalTickets:     500,
			AvailableTickets: 320,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			ID:               "evt-002",
			Title:            "Indie Music Festival",
			Description:      "A festival featuring top indie bands from across Southeast Asia.",
			EventDate:        time.Date(2025, 5, 21, 18, 30, 0, 0, time.Local),
			Venue:            "Lapangan D Senayan",
			TicketPrice:      150000,
			TotalTickets:     1000,
			AvailableTickets: 750,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			ID:               "evt-003",
			Title:            "Startup Pitch Day",
			Description:      "Pitch event for early-stage startups to meet investors and mentors.",
			EventDate:        time.Date(2025, 7, 10, 10, 0, 0, 0, time.Local),
			Venue:            "GoWork - Kemang",
			TicketPrice:      100000,
			TotalTickets:     200,
			AvailableTickets: 120,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	return MockEvents, nil
}

func (r *eventRepository) Create(ctx context.Context, event *models.Event) error {
	// TODO: implement event creation
	// Example: INSERT INTO events (id, title, description, event_date, ...) VALUES (...)
	return fmt.Errorf("not implemented")
}

func (r *eventRepository) Update(ctx context.Context, event *models.Event) error {
	// TODO: implement event update
	// Example: UPDATE events SET title = $1, description = $2, ... WHERE id = $n
	return fmt.Errorf("not implemented")
}

func (r *eventRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement event deletion
	// Example: DELETE FROM events WHERE id = $1
	return fmt.Errorf("not implemented")
}

func (r *eventRepository) DecrementAvailableTickets(ctx context.Context, eventID string, quantity int) error {
	// TODO: implement ticket inventory decrement with row-level locking
	// Example: UPDATE events SET available_tickets = available_tickets - $1 WHERE id = $2 AND available_tickets >= $1
	// Important: Use FOR UPDATE lock in transaction to prevent double-booking
	return fmt.Errorf("not implemented")
}
