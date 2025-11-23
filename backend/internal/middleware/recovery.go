package middleware

import (
	"net/http"

	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RecoveryMiddleware(log zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().
					Interface("error", err).
					Str("path", c.Request.URL.Path).
					Msg("panic recovered")

				response.Error(c, http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()

		c.Next()
	}
}