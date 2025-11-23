package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func setupEventRoutes(rg *gin.RouterGroup, h *handlers.EventHandler) {
	events := rg.Group("/events")
	{
		// Public routes
		events.GET("", h.List)
		events.GET("/:id", h.GetByID)

		// Protected routes (admin only)
		events.POST("", middleware.AuthMiddleware(), h.Create)
		events.PUT("/:id", middleware.AuthMiddleware(), h.Update)
		events.DELETE("/:id", middleware.AuthMiddleware(), h.Delete)
	}
}