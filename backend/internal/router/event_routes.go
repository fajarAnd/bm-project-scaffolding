package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func setupEventRoutes(rg *gin.RouterGroup, h *handlers.EventHandler, jwtCfg config.JWTConfig) {
	events := rg.Group("/events")
	{
		// Public routes
		events.GET("", h.List)
		events.GET("/:id", h.GetByID)

		// Protected routes (admin only)
		events.POST("", middleware.AuthMiddleware(jwtCfg), middleware.RequireRole(models.RoleAdmin), h.Create)
		events.PUT("/:id", middleware.AuthMiddleware(jwtCfg), middleware.RequireRole(models.RoleAdmin), h.Update)
		events.DELETE("/:id", middleware.AuthMiddleware(jwtCfg), middleware.RequireRole(models.RoleAdmin), h.Delete)
	}
}