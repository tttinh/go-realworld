package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getArticle(c *gin.Context) {
	userID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, userID, slug)
	if err != nil {
		abort(c, err)
		return
	}

	var res articleRes
	res.fromEntity(a)
	c.JSON(http.StatusOK, res)
}
