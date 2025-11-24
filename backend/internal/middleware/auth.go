package middleware

import (
	"net/http"
	"strings"

	"github.com/baramulti/ticketing-system/backend/internal/config"
	jwtutil "github.com/baramulti/ticketing-system/backend/pkg/jwt"
	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	AuthHeaderKey  = "Authorization"
	UserContextKey = "user"
	UserIDKey      = "user_id"
	UserEmailKey   = "user_email"
	UserRolesKey   = "user_roles"
)

func AuthMiddleware(jwtCfg config.JWTConfig) gin.HandlerFunc {
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

		tokenString := parts[1]

		// Validate JWT token
		claims, err := jwtutil.ValidateToken(tokenString, jwtCfg.Secret)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		// Set user info in context
		c.Set(UserIDKey, claims.UserID)
		c.Set(UserEmailKey, claims.Email)
		c.Set(UserRolesKey, claims.Roles)

		c.Next()
	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesVal, exists := c.Get(UserRolesKey)
		if !exists {
			response.Error(c, http.StatusUnauthorized, "user not authenticated")
			c.Abort()
			return
		}

		userRoles, ok := rolesVal.([]string)
		if !ok {
			response.Error(c, http.StatusInternalServerError, "invalid roles")
			c.Abort()
			return
		}
		
		for _, userRole := range userRoles {
			for _, allowed := range allowedRoles {
				if userRole == allowed {
					c.Next()
					return
				}
			}
		}

		response.Error(c, http.StatusForbidden, "insufficient permissions")
		c.Abort()
	}
}
