package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseFeed(c *gin.Context) {
	var res batchArticleRes
	c.JSON(http.StatusOK, res)
}
