package services

import (
	"context"
	"fmt"

	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/rs/zerolog"
)

type EventService interface {
	GetByID(ctx context.Context, id string) (*models.Event, error)
	List(ctx context.Context, page, pageSize int) ([]*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
	Update(ctx context.Context, event *models.Event) error
	Delete(ctx context.Context, id string) error
}

type eventService struct {
	repo repositories.EventRepository
	log  zerolog.Logger
}

func NewEventService(repo repositories.EventRepository, log zerolog.Logger) EventService {
	return &eventService{
		repo: repo,
		log:  log,
	}
}

func (s *eventService) GetByID(ctx context.Context, id string) (*models.Event, error) {
	// TODO: fetch event with validation
	return s.repo.FindByID(ctx, id)
}

func (s *eventService) List(ctx context.Context, page, pageSize int) ([]*models.Event, error) {
	offset := (page - 1) * pageSize
	// TODO: add caching layer here (Redis)
	return s.repo.List(ctx, pageSize, offset)
}

func (s *eventService) Create(ctx context.Context, event *models.Event) error {
	// TODO: validate and create event
	// set timestamps, generate ID, etc.
	return fmt.Errorf("not implemented")
}

func (s *eventService) Update(ctx context.Context, event *models.Event) error {
	// TODO: validate event exists before update
	return fmt.Errorf("not implemented")
}

func (s *eventService) Delete(ctx context.Context, id string) error {
	// TODO: check for active tickets before deletion
	return fmt.Errorf("not implemented")
}