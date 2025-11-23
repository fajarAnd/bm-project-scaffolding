package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func setupAuthRoutes(rg *gin.RouterGroup, h *handlers.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		// TODO: add /logout, /refresh-token endpoints
	}
}