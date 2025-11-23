package services

import (
	"context"
	"fmt"

	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/rs/zerolog"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)
	ValidateToken(ctx context.Context, token string) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
	log      zerolog.Logger
}

func NewAuthService(userRepo repositories.UserRepository, log zerolog.Logger) AuthService {
	return &authService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// TODO: implement login flow
	// - find user by email
	// - compare password hash (bcrypt)
	// - generate JWT token
	s.log.Info().Str("email", req.Email).Msg("login attempt")
	return nil, fmt.Errorf("not implemented")
}

func (s *authService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// TODO: implement registration
	// - check email uniqueness
	// - hash password
	// - create user
	// - generate token
	s.log.Info().Str("email", req.Email).Msg("registration attempt")
	return nil, fmt.Errorf("not implemented")
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	// TODO: parse and validate JWT
	// - verify signature
	// - check expiry
	// - extract user ID
	// - fetch user details
	return nil, fmt.Errorf("not implemented")
}