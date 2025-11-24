package router

import (
	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type RouterConfig struct {
	Config        *config.Config
	Logger        zerolog.Logger
	AuthHandler   *handlers.AuthHandler
	EventHandler  *handlers.EventHandler
	TicketHandler *handlers.TicketHandler
	UserHandler   *handlers.UserHandler
}

func Setup(cfg *RouterConfig) *gin.Engine {
	if cfg.Config.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global middleware
	r.Use(middleware.RecoveryMiddleware(cfg.Logger))
	r.Use(middleware.LoggingMiddleware(cfg.Logger))
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.HEAD("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	api := r.Group("/api/v1")
	{
		setupAuthRoutes(api, cfg.AuthHandler)
		setupEventRoutes(api, cfg.EventHandler, cfg.Config.JWT)
		setupTicketRoutes(api, cfg.TicketHandler, cfg.Config.JWT)
		setupUserRoutes(api, cfg.UserHandler, cfg.Config.JWT)
	}

	return r
}
