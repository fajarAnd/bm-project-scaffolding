package services

import (
	"context"
	"fmt"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	jwtutil "github.com/baramulti/ticketing-system/backend/pkg/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)
	ValidateToken(ctx context.Context, token string) (*models.User, error)
}

type authService struct {
	userRepo  repositories.UserRepository
	jwtConfig config.JWTConfig
	log       zerolog.Logger
}

func NewAuthService(userRepo repositories.UserRepository, jwtConfig config.JWTConfig, log zerolog.Logger) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
		log:       log,
	}
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// TODO: Production flow:
	// 1. Find user by email: user, err := s.userRepo.FindByEmail(ctx, req.Email)
	// 2. Compare password: bcrypt.CompareHashAndPassword(user.PasswordHash, req.Password)
	// 3. Load user roles from database

	s.log.Info().Str("email", req.Email).Msg("login attempt (stub)")

	// STUB: Mock user for development
	mockUser := &models.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock role - check if admin email
	roles := []string{models.RoleUser}
	if req.Email == "admin@example.com" {
		roles = []string{models.RoleAdmin}
	}

	// Generate real JWT token
	expiry, _ := time.ParseDuration(s.jwtConfig.Expiry)
	token, err := jwtutil.GenerateToken(mockUser.ID, mockUser.Email, roles, s.jwtConfig.Secret, expiry)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to generate token")
		return nil, fmt.Errorf("failed to generate token")
	}

	return &dto.AuthResponse{
		Token: token,
		User:  mockUser,
	}, nil
}

func (s *authService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// TODO: Production flow:
	// 1. Check email uniqueness
	// 2. Hash password with bcrypt
	// 3. Create user in database
	// 4. Assign default role

	s.log.Info().Str("email", req.Email).Msg("registration attempt (stub)")

	// STUB: Mock new user
	newUser := &models.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Default role is "user"
	role := models.RoleUser
	if req.Role != "" {
		role = req.Role
	}

	// Generate JWT
	expiry, _ := time.ParseDuration(s.jwtConfig.Expiry)
	token, err := jwtutil.GenerateToken(newUser.ID, newUser.Email, []string{role}, s.jwtConfig.Secret, expiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &dto.AuthResponse{
		Token: token,
		User:  newUser,
	}, nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	// Validate JWT signature and expiry
	claims, err := jwtutil.ValidateToken(token, s.jwtConfig.Secret)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// TODO: Fetch user from database to check if still active
	// user, err := s.userRepo.FindByID(ctx, claims.UserID)

	// STUB: Return mock user from claims
	user := &models.User{
		ID:        claims.UserID,
		Email:     claims.Email,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}