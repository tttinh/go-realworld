package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
