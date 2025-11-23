package middleware

import (
	"net/http"
	"strings"

	"github.com/baramulti/ticketing-system/backend/internal/models"
	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	AuthHeaderKey  = "Authorization"
	UserContextKey = "user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthHeaderKey)
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header format")
			c.Abort()
			return
		}

		_ = parts[1] // token - TODO: validate JWT

		// TODO: validate JWT token and extract user info
		// For now, mock user for development
		mockUser := &models.User{
			// ID:    "mock-user-id",
			// Email: "demo@example.com",
			// Role:  models.RoleUser,
		}

		c.Set(UserContextKey, mockUser)
		c.Next()
	}
}

func RequireRole(allowedRoles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userVal, exists := c.Get(UserContextKey)
		if !exists {
			response.Error(c, http.StatusUnauthorized, "user not authenticated")
			c.Abort()
			return
		}

		_, ok := userVal.(*models.User)
		if !ok {
			response.Error(c, http.StatusInternalServerError, "invalid user context")
			c.Abort()
			return
		}

		// TODO: check if user.Role is in allowedRoles
		// For now, allow all authenticated users
		c.Next()
	}
}
