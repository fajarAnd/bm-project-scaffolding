package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Summary: Tests ticket purchase with various quantities and scenarios
func TestTicketService_PurchaseTicket(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		eventID  string
		quantity int
		wantErr  bool
	}{
		{
			name:     "purchase single ticket",
			userID:   "user-001",
			eventID:  "event-001",
			quantity: 1,
			wantErr:  false,
		},
		{
			name:     "purchase multiple tickets",
			userID:   "user-002",
			eventID:  "event-002",
			quantity: 5,
			wantErr:  false,
		},
		{
			name:     "purchase max tickets",
			userID:   "user-003",
			eventID:  "event-003",
			quantity: 10,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTicketRepo := mocks.NewTicketRepository(t)
			mockEventRepo := mocks.NewEventRepository(t)
			logger := zerolog.Nop()

			service := NewTicketService(mockTicketRepo, mockEventRepo, logger)

			req := &dto.PurchaseRequest{
				EventID:  tt.eventID,
				Quantity: tt.quantity,
			}

			resp, err := service.PurchaseTicket(context.Background(), tt.userID, req)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.OrderID)
			assert.NotEmpty(t, resp.TransactionID)
			assert.Equal(t, "confirmed", resp.Status)
			assert.Contains(t, resp.Message, fmt.Sprintf("%d ticket(s)", tt.quantity))
		})
	}
}

// TestTicketService_GetUserOrders
// Purpose: Verify repository integration for listing orders
func TestTicketService_GetUserOrders(t *testing.T) {
	mockTicketRepo := mocks.NewTicketRepository(t)
	mockEventRepo := mocks.NewEventRepository(t)
	logger := zerolog.Nop()

	tests := []struct {
		name           string
		userID         string
		mockOrders     []*models.TicketOrder
		mockErr        error
		expectedCount  int
		expectError    bool
	}{
		{
			name:   "user with multiple orders",
			userID: "user-123",
			mockOrders: []*models.TicketOrder{
				{
					ID:         "order-1",
					EventID:    "event-1",
					UserID:     "user-123",
					Quantity:   2,
					TotalPrice: 500000,
					Status:     "confirmed",
					CreatedAt:  time.Now(),
				},
				{
					ID:         "order-2",
					EventID:    "event-2",
					UserID:     "user-123",
					Quantity:   1,
					TotalPrice: 250000,
					Status:     "confirmed",
					CreatedAt:  time.Now(),
				},
			},
			mockErr:       nil,
			expectedCount: 2,
			expectError:   false,
		},
		{
			name:          "user with no orders",
			userID:        "user-456",
			mockOrders:    []*models.TicketOrder{},
			mockErr:       nil,
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:          "repository error",
			userID:        "user-789",
			mockOrders:    nil,
			mockErr:       fmt.Errorf("database connection error"),
			expectedCount: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTicketRepo.On("ListOrdersByUserID", mock.Anything, tt.userID).
				Return(tt.mockOrders, tt.mockErr).
				Once()

			service := NewTicketService(mockTicketRepo, mockEventRepo, logger)
			orders, err := service.GetUserOrders(context.Background(), tt.userID)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, orders, tt.expectedCount)

			if tt.expectedCount > 0 {
				for i, order := range orders {
					assert.Equal(t, tt.mockOrders[i].ID, order.ID)
					assert.Equal(t, tt.userID, order.UserID)
				}
			}

			mockTicketRepo.AssertExpectations(t)
		})
	}
}

// Purpose: Validate single order retrieval from repository
func TestTicketService_GetOrderByID(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		orderID     string
		mockOrder   *models.TicketOrder
		mockErr     error
		expectError bool
	}{
		{
			name:    "valid order found",
			orderID: "order-valid-123",
			mockOrder: &models.TicketOrder{
				ID:         "order-valid-123",
				EventID:    "event-001",
				UserID:     "user-001",
				Quantity:   3,
				TotalPrice: 750000,
				Status:     "confirmed",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			mockErr:     nil,
			expectError: false,
		},
		{
			name:        "order not found",
			orderID:     "order-nonexistent",
			mockOrder:   nil,
			mockErr:     fmt.Errorf("order not found"),
			expectError: true,
		},
		{
			name:    "order with pending status",
			orderID: "order-pending-456",
			mockOrder: &models.TicketOrder{
				ID:         "order-pending-456",
				EventID:    "event-002",
				UserID:     "user-002",
				Quantity:   1,
				TotalPrice: 150000,
				Status:     "pending",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			mockErr:     nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTicketRepo := mocks.NewTicketRepository(t)
			mockEventRepo := mocks.NewEventRepository(t)
			logger := zerolog.Nop()

			mockTicketRepo.On("FindOrderByID", mock.Anything, tt.orderID).
				Return(tt.mockOrder, tt.mockErr).
				Once()

			service := NewTicketService(mockTicketRepo, mockEventRepo, logger)
			order, err := service.GetOrderByID(context.Background(), tt.orderID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, order)
				mockTicketRepo.AssertExpectations(t)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, order)
			assert.Equal(t, tt.mockOrder.ID, order.ID)
			assert.Equal(t, tt.mockOrder.UserID, order.UserID)
			assert.Equal(t, tt.mockOrder.Quantity, order.Quantity)
			assert.Equal(t, tt.mockOrder.Status, order.Status)

			mockTicketRepo.AssertExpectations(t)
		})
	}
}

// Purpose: Ensure response fields are properly populated
func TestTicketService_PurchaseTicket_ResponseFormat(t *testing.T) {
	mockTicketRepo := mocks.NewTicketRepository(t)
	mockEventRepo := mocks.NewEventRepository(t)
	logger := zerolog.Nop()

	service := NewTicketService(mockTicketRepo, mockEventRepo, logger)

	req := &dto.PurchaseRequest{
		EventID:  "event-test",
		Quantity: 2,
	}

	resp, err := service.PurchaseTicket(context.Background(), "user-test", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Check all response fields are populated
	assert.NotEmpty(t, resp.OrderID, "OrderID should not be empty")
	assert.NotEmpty(t, resp.TransactionID, "TransactionID should not be empty")
	assert.NotEmpty(t, resp.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.Message, "Message should not be empty")

	// Verify transaction ID format
	assert.Contains(t, resp.TransactionID, "TXN-", "TransactionID should contain TXN- prefix")
}