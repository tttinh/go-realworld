package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
