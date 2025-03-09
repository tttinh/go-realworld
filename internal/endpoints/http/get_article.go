package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getArticle(c *gin.Context) {
	viewerID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, viewerID, slug)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var res articleRes
	res.fromEntity(a)
	c.JSON(http.StatusOK, res)
}
