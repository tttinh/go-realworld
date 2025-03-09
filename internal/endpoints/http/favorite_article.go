package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) favoriteArticle(c *gin.Context) {
	authorID, _ := h.jwt.GetUserID(c)
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.articles.AddFavorite(c, authorID, a.ID); err != nil {
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
