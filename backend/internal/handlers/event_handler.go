package handlers

import (
	"net/http"
	"strconv"

	"github.com/baramulti/ticketing-system/backend/internal/services"
	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventSvc services.EventService
}

func NewEventHandler(eventSvc services.EventService) *EventHandler {
	return &EventHandler{eventSvc: eventSvc}
}

func (h *EventHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	event, err := h.eventSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "event not found")
		return
	}

	response.Success(c, http.StatusOK, event)
}

func (h *EventHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	events, err := h.eventSvc.List(c.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"events": events,
		"page":   page,
	})
}

func (h *EventHandler) Create(c *gin.Context) {
	// TODO: parse CreateEventRequest DTO
	// TODO: validate admin role
	response.Error(c, http.StatusNotImplemented, "not implemented")
}

func (h *EventHandler) Update(c *gin.Context) {
	// TODO: parse UpdateEventRequest DTO
	// TODO: validate admin role
	response.Error(c, http.StatusNotImplemented, "not implemented")
}

func (h *EventHandler) Delete(c *gin.Context) {
	// TODO: validate admin role
	response.Error(c, http.StatusNotImplemented, "not implemented")
}