package services

import (
	"context"
	"testing"
	"time"

	"github.com/baramulti/ticketing-system/backend/internal/config"
	"github.com/baramulti/ticketing-system/backend/internal/dto"
	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/internal/repositories/mocks"
	jwtutil "github.com/baramulti/ticketing-system/backend/pkg/jwt"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestAuthService_Login
// Summary: Tests the Login method with different user scenarios
// Purpose: Validate JWT token generation and user role assignment in stub mode
func TestAuthService_Login(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr bool
		checkRole   string
	}{
		{
			name:        "normal user login",
			email:       "user@example.com",
			password:    "password123",
			expectedErr: false,
			checkRole:   models.RoleUser,
		},
		{
			name:        "admin user login",
			email:       "admin@example.com",
			password:    "admin123",
			expectedErr: false,
			checkRole:   models.RoleAdmin,
		},
		{
			name:        "another regular user",
			email:       "test@example.com",
			password:    "test123",
			expectedErr: false,
			checkRole:   models.RoleUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mocks.NewUserRepository(t)
			logger := zerolog.Nop()
			jwtConfig := config.JWTConfig{
				Secret: "test-secret-key",
				Expiry: "24h",
			}

			service := NewAuthService(mockUserRepo, jwtConfig, logger)

			req := &dto.LoginRequest{
				Email:    tt.email,
				Password: tt.password,
			}

			resp, err := service.Login(context.Background(), req)

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Token)
			assert.NotNil(t, resp.User)
			assert.Equal(t, tt.email, resp.User.Email)
			assert.True(t, resp.User.IsActive)

			// Verify token is valid
			claims, err := jwtutil.ValidateToken(resp.Token, jwtConfig.Secret)
			assert.NoError(t, err)
			assert.Equal(t, tt.email, claims.Email)
			assert.Contains(t, claims.Roles, tt.checkRole)
		})
	}
}

// TestAuthService_Register
// Summary: Tests user registration with different role assignments
// Purpose: Ensure proper user creation and token generation in stub mode
func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		password     string
		role         string
		expectedRole string
		expectedErr  bool
	}{
		{
			name:         "register default user",
			email:        "newuser@example.com",
			password:     "password123",
			role:         "",
			expectedRole: models.RoleUser,
			expectedErr:  false,
		},
		{
			name:         "register with admin role",
			email:        "newadmin@example.com",
			password:     "admin123",
			role:         models.RoleAdmin,
			expectedRole: models.RoleAdmin,
			expectedErr:  false,
		},
		{
			name:         "register organizer",
			email:        "organizer@example.com",
			password:     "org123",
			role:         models.RoleOrganizer,
			expectedRole: models.RoleOrganizer,
			expectedErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mocks.NewUserRepository(t)
			logger := zerolog.Nop()
			jwtConfig := config.JWTConfig{
				Secret: "test-secret",
				Expiry: "24h",
			}

			service := NewAuthService(mockUserRepo, jwtConfig, logger)

			req := &dto.RegisterRequest{
				Email:    tt.email,
				Password: tt.password,
				Role:     tt.role,
			}

			resp, err := service.Register(context.Background(), req)

			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Token)
			assert.Equal(t, tt.email, resp.User.Email)
			assert.True(t, resp.User.IsActive)

			// Check JWT claims
			claims, err := jwtutil.ValidateToken(resp.Token, jwtConfig.Secret)
			assert.NoError(t, err)
			assert.Equal(t, tt.email, claims.Email)
			assert.Contains(t, claims.Roles, tt.expectedRole)
		})
	}
}

// TestAuthService_ValidateToken
// Summary: Tests token validation with various token states
// Purpose: Verify JWT validation logic and user reconstruction from claims
func TestAuthService_ValidateToken(t *testing.T) {
	logger := zerolog.Nop()
	jwtConfig := config.JWTConfig{
		Secret: "test-secret-key-validate",
		Expiry: "1h",
	}

	tests := []struct {
		name        string
		setupToken  func() string
		expectError bool
		checkEmail  string
	}{
		{
			name: "valid token",
			setupToken: func() string {
				expiry, _ := time.ParseDuration(jwtConfig.Expiry)
				token, _ := jwtutil.GenerateToken(
					"user-123",
					"valid@example.com",
					[]string{models.RoleUser},
					jwtConfig.Secret,
					expiry,
				)
				return token
			},
			expectError: false,
			checkEmail:  "valid@example.com",
		},
		{
			name: "valid admin token",
			setupToken: func() string {
				expiry, _ := time.ParseDuration(jwtConfig.Expiry)
				token, _ := jwtutil.GenerateToken(
					"admin-456",
					"admin@example.com",
					[]string{models.RoleAdmin},
					jwtConfig.Secret,
					expiry,
				)
				return token
			},
			expectError: false,
			checkEmail:  "admin@example.com",
		},
		{
			name: "invalid token - wrong secret",
			setupToken: func() string {
				expiry, _ := time.ParseDuration(jwtConfig.Expiry)
				token, _ := jwtutil.GenerateToken(
					"user-789",
					"wrong@example.com",
					[]string{models.RoleUser},
					"wrong-secret",
					expiry,
				)
				return token
			},
			expectError: true,
		},
		{
			name: "invalid token - malformed",
			setupToken: func() string {
				return "this-is-not-a-valid-jwt-token"
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mocks.NewUserRepository(t)
			service := NewAuthService(mockUserRepo, jwtConfig, logger)

			token := tt.setupToken()
			user, err := service.ValidateToken(context.Background(), token)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, tt.checkEmail, user.Email)
			assert.True(t, user.IsActive)
		})
	}
}

// TestAuthService_TokenExpiry
// Summary: Validates that expired tokens are properly rejected
// Purpose: Test JWT expiration handling
func TestAuthService_TokenExpiry(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	logger := zerolog.Nop()
	jwtConfig := config.JWTConfig{
		Secret: "test-expiry-secret",
		Expiry: "1ns", // very short expiry
	}

	service := NewAuthService(mockUserRepo, jwtConfig, logger)

	// Generate token that will expire immediately
	expiry, _ := time.ParseDuration(jwtConfig.Expiry)
	token, err := jwtutil.GenerateToken(
		"user-expired",
		"expired@example.com",
		[]string{models.RoleUser},
		jwtConfig.Secret,
		expiry,
	)
	assert.NoError(t, err)

	// Wait a bit to ensure token expires
	time.Sleep(10 * time.Millisecond)

	user, err := service.ValidateToken(context.Background(), token)
	assert.Error(t, err)
	assert.Nil(t, user)
}