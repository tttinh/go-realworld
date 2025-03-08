package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
