package dto

import "github.com/baramulti/ticketing-system/backend/internal/models"

type PurchaseRequest struct {
	EventID  string `json:"event_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,min=1,max=10"`
}

type PurchaseResponse struct {
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"` // from payment gateway
	Status        string `json:"status"`
	Message       string `json:"message,omitempty"`
}

type OrderResponse struct {
	Order   *models.TicketOrder `json:"order"`
	Tickets []*models.Ticket    `json:"tickets,omitempty"`
}

type OrderListResponse struct {
	Orders []*models.TicketOrder `json:"orders"`
	Total  int                   `json:"total"`
}