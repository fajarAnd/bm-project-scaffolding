package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func setupTicketRoutes(rg *gin.RouterGroup, h *handlers.TicketHandler) {
	tickets := rg.Group("/tickets")
	tickets.Use(middleware.AuthMiddleware()) // All ticket routes require auth
	{
		tickets.POST("/purchase", h.Purchase)
		tickets.GET("/my-orders", h.GetUserOrders)
		// TODO: add /orders/:id for order details
	}
}