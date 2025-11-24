package handlers

import (
	"net/http"

	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/services"
	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketSvc services.TicketService
}

func NewTicketHandler(ticketSvc services.TicketService) *TicketHandler {
	return &TicketHandler{ticketSvc: ticketSvc}
}

func (h *TicketHandler) Purchase(c *gin.Context) {
	// Extract user ID from JWT context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "user not authenticated")
		return
	}

	var req dto.PurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	result, err := h.ticketSvc.PurchaseTicket(c.Request.Context(), userID.(string), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (h *TicketHandler) GetUserOrders(c *gin.Context) {
	// Extract user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "user not authenticated")
		return
	}

	orders, err := h.ticketSvc.GetUserOrders(c.Request.Context(), userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"orders": orders})
}