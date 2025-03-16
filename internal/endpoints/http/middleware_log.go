package httpendpoints

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		c.Next()

		// Stop timer
		l := log.With(
			"method", c.Request.Method,
			"path", path,
			"elapsed", time.Since(start),
			"status", c.Writer.Status(),
			"client_ip", c.ClientIP(),
			"response_size", c.Writer.Size(),
		)

		if len(c.Errors) == 0 {
			l.Info("http")
			return
		}

		// Client errors
		if c.Writer.Status() < http.StatusInternalServerError {
			l.Warn("http", "err", errors.Unwrap(c.Errors.Last()))
			return
		}

		// Server errors
		l.Error("http", "err", errors.Unwrap(c.Errors.Last()))
	}
}
