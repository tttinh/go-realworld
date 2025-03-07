package httpendpoints

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getArticle(c *gin.Context) {
	viewerID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, viewerID, slug)
	if err != nil {
		error400(c, err)
		return
	}

	var res articleRes
	res.fromEntity(a)
	ok(c, res)
}
