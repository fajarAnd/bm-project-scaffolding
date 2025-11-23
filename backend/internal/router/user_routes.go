package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func setupUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler) {
	users := rg.Group("/users")
	users.Use(middleware.AuthMiddleware()) // All user routes require auth
	{
		users.GET("/me", h.GetMe)
		users.PUT("/me", h.Update)

		// Admin only routes
		// TODO: add RequireRole middleware
		users.POST("", h.Create)
		users.DELETE("/:id", h.Delete)
	}
}