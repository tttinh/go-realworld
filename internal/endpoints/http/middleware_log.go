package httpendpoints

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LogMiddleware() gin.HandlerFunc {
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
		zlog := log.Info()
		if len(c.Errors) > 0 {
			zlog = log.Error()
		}

		zlog.
			Str("method", c.Request.Method).
			Str("path", path).
			Dur("elapsed", time.Since(start)).
			Int("status", c.Writer.Status()).
			Str("client_ip", c.ClientIP()).
			Int("response_size", c.Writer.Size()).
			Err(c.Errors.Last()).
			Send()
	}
}
