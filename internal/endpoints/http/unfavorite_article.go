package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) unfavoriteArticle(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.articles.RemoveFavorite(c, authorID, a.ID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, a.Slug)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}
