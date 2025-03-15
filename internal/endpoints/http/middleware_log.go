package httpendpoints

import (
	"log/slog"
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

		if len(c.Errors) > 0 {
			l.Error("http", "err", c.Errors.Last())
		} else {
			l.Info("http")
		}
	}
}
