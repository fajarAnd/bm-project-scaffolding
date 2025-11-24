package services

import (
	"context"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type UserService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, email, password string) (*models.User, error)
	Update(ctx context.Context, id string, email string) error
	Delete(ctx context.Context, id string) error
	GetUserRoles(ctx context.Context, userID string) ([]*models.Role, error)
}

type userService struct {
	userRepo repositories.UserRepository
	log      zerolog.Logger
}

func NewUserService(userRepo repositories.UserRepository, log zerolog.Logger) UserService {
	return &userService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	// STUB: return mock user for now
	// TODO: call s.userRepo.FindByID(ctx, id) when repository is implemented
	s.log.Debug().Str("user_id", id).Msg("fetching user by id")

	return &models.User{
		ID:        id,
		Email:     "stub@example.com",
		IsActive:  true,
		CreatedAt: time.Now().Add(-30 * 24 * time.Hour),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	// STUB: return mock user
	s.log.Debug().Str("email", email).Msg("fetching user by email")

	return &models.User{
		ID:        "a1b2c3d4-5678-90ab-cdef-1234567890ab",
		Email:     email,
		IsActive:  true,
		CreatedAt: time.Now().Add(-30 * 24 * time.Hour),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *userService) Create(ctx context.Context, email, password string) (*models.User, error) {
	// TODO: hash password with bcrypt
	// TODO: validate email format
	// TODO: check email uniqueness
	// TODO: call s.userRepo.Create()

	s.log.Info().Str("email", email).Msg("creating new user (stub)")

	newUser := &models.User{
		ID:        uuid.New().String(),
		Email:     email,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return newUser, nil
}

func (s *userService) Update(ctx context.Context, id string, email string) error {
	// TODO: validate email format
	// TODO: check email uniqueness (if changed)
	// TODO: update in database
	s.log.Info().Str("user_id", id).Str("email", email).Msg("updating user")
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	// TODO: check if user has active tickets/orders
	// TODO: implement soft delete (set deleted_at)
	s.log.Info().Str("user_id", id).Msg("deleting user")
	return nil
}

func (s *userService) GetUserRoles(ctx context.Context, userID string) ([]*models.Role, error) {
	// STUB: return mock roles
	// Production: query from user_roles + roles tables
	s.log.Debug().Str("user_id", userID).Msg("fetching user roles")

	// Mock: return "user" role for regular users
	return []*models.Role{
		{
			ID:          uuid.New().String(),
			Name:        models.RoleUser,
			Description: "Regular user",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}