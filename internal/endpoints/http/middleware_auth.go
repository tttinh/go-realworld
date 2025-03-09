package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwt JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.Validate(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newErrorRes(err))
		}

		c.Next()
	}
}
