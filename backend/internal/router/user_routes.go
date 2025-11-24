package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func setupUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler, jwtCfg config.JWTConfig) {
	users := rg.Group("/users")
	users.Use(middleware.AuthMiddleware(jwtCfg)) // All user routes require auth
	{
		users.GET("/me", h.GetMe)
		users.PUT("/me", h.Update)

		// Admin only routes
		users.POST("", middleware.RequireRole(models.RoleAdmin), h.Create)
		users.DELETE("/:id", middleware.RequireRole(models.RoleAdmin), h.Delete)
	}
}