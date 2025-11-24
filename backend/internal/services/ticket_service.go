package services

import (
	"context"
	"fmt"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/google/uuid"
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
	// TODO: Production flow:
	// 1. Validate event exists: event, err := s.eventRepo.FindByID(ctx, req.EventID)
	// 2. Check availability: event.AvailableTickets >= req.Quantity
	// 3. Start database transaction
	// 4. Lock event row: SELECT FOR UPDATE
	// 5. Process payment via gateway (Stripe/Midtrans)
	// 6. Create order record with status "pending"
	// 7. Decrement event.AvailableTickets
	// 8. Generate unique ticket codes
	// 9. Update order status to "confirmed"
	// 10. Commit transaction
	// 11. Send confirmation email (async job)

	s.log.Info().
		Str("user_id", userID).
		Str("event_id", req.EventID).
		Int("qty", req.Quantity).
		Msg("purchase attempt (stub)")

	// STUB: Return mock successful purchase
	orderID := uuid.New().String()
	transactionID := fmt.Sprintf("TXN-%d", time.Now().Unix())

	s.log.Info().
		Str("order_id", orderID).
		Str("transaction_id", transactionID).
		Msg("ticket purchase successful (mock)")

	return &dto.PurchaseResponse{
		OrderID:       orderID,
		TransactionID: transactionID,
		Status:        "confirmed",
		Message:       fmt.Sprintf("Successfully purchased %d ticket(s)", req.Quantity),
	}, nil
}

func (s *ticketService) GetUserOrders(ctx context.Context, userID string) ([]*models.TicketOrder, error) {
	// TODO: fetch user's orders with event details
	return s.ticketRepo.ListOrdersByUserID(ctx, userID)
}

func (s *ticketService) GetOrderByID(ctx context.Context, orderID string) (*models.TicketOrder, error) {
	// TODO: fetch order with tickets
	return s.ticketRepo.FindOrderByID(ctx, orderID)
}