package main

import (
	"fmt"
	"os"

	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/handlers"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/baramulti/ticketing-system/backend/internal/router"
	"github.com/baramulti/ticketing-system/backend/internal/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// Setup logger
	logger := setupLogger(cfg.Server.Env)

	// Connect to database
	db, err := connectDB(cfg.Database.URL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer db.Close()

	logger.Info().Msg("database connected")

	// Initialize dependencies
	repos := initRepositories(db)
	services := initServices(repos, logger)
	handlers := initHandlers(services)

	// Setup router
	r := router.Setup(&router.RouterConfig{
		Config:        cfg,
		Logger:        logger,
		AuthHandler:   handlers.auth,
		EventHandler:  handlers.event,
		TicketHandler: handlers.ticket,
		UserHandler:   handlers.user,
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	logger.Info().Str("addr", addr).Str("env", cfg.Server.Env).Msg("server starting")

	if err := r.Run(addr); err != nil {
		logger.Fatal().Err(err).Msg("server failed to start")
	}
}

func setupLogger(env string) zerolog.Logger {
	if env == "development" {
		return log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	return log.Logger
}

func connectDB(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	return db, nil

	//return nil, fmt.Errorf("database connection not implemented")
}

type repositoryDeps struct {
	user   repositories.UserRepository
	event  repositories.EventRepository
	ticket repositories.TicketRepository
}

func initRepositories(db *sqlx.DB) *repositoryDeps {
	return &repositoryDeps{
		user:   repositories.NewUserRepository(db),
		event:  repositories.NewEventRepository(db),
		ticket: repositories.NewTicketRepository(db),
	}
}

type serviceDeps struct {
	auth   services.AuthService
	event  services.EventService
	ticket services.TicketService
}

func initServices(repos *repositoryDeps, logger zerolog.Logger) *serviceDeps {
	return &serviceDeps{
		auth:   services.NewAuthService(repos.user, logger),
		event:  services.NewEventService(repos.event, logger),
		ticket: services.NewTicketService(repos.ticket, repos.event, logger),
	}
}

type handlerDeps struct {
	auth   *handlers.AuthHandler
	event  *handlers.EventHandler
	ticket *handlers.TicketHandler
	user   *handlers.UserHandler
}

func initHandlers(services *serviceDeps) *handlerDeps {
	return &handlerDeps{
		auth:   handlers.NewAuthHandler(services.auth),
		event:  handlers.NewEventHandler(services.event),
		ticket: handlers.NewTicketHandler(services.ticket),
		user:   handlers.NewUserHandler(),
	}
}
