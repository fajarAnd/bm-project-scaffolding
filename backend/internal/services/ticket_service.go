package services

import (
	"context"
	"fmt"

	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/rs/zerolog"
)

type TicketService interface {
	PurchaseTicket(ctx context.Context, userID string, req *dto.PurchaseRequest) (*dto.PurchaseResponse, error)
	GetUserOrders(ctx context.Context, userID string) ([]*models.TicketOrder, error)
	GetOrderByID(ctx context.Context, orderID string) (*models.TicketOrder, error)
}

type ticketService struct {
	ticketRepo repositories.TicketRepository
	eventRepo  repositories.EventRepository
	log        zerolog.Logger
}

func NewTicketService(
	ticketRepo repositories.TicketRepository,
	eventRepo repositories.EventRepository,
	log zerolog.Logger,
) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		eventRepo:  eventRepo,
		log:        log,
	}
}

func (s *ticketService) PurchaseTicket(ctx context.Context, userID string, req *dto.PurchaseRequest) (*dto.PurchaseResponse, error) {
	// TODO: implement purchase flow with transaction
	// 1. check event availability
	// 2. start transaction
	// 3. lock event row (SELECT FOR UPDATE)
	// 4. process payment (stub)
	// 5. create order
	// 6. decrement tickets
	// 7. generate ticket codes
	// 8. commit
	// 9. send email (async)

	s.log.Info().
		Str("user_id", userID).
		Str("event_id", req.EventID).
		Int("qty", req.Quantity).
		Msg("purchase attempt")

	return nil, fmt.Errorf("not implemented")
}

func (s *ticketService) GetUserOrders(ctx context.Context, userID string) ([]*models.TicketOrder, error) {
	// TODO: fetch user's orders with event details
	return s.ticketRepo.ListOrdersByUserID(ctx, userID)
}

func (s *ticketService) GetOrderByID(ctx context.Context, orderID string) (*models.TicketOrder, error) {
	// TODO: fetch order with tickets
	return s.ticketRepo.FindOrderByID(ctx, orderID)
}