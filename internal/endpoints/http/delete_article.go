package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteArticle(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if authorID != a.AuthorID {
		c.AbortWithError(http.StatusForbidden, ErrAccessForbidden)
		return
	}

	if err := h.articles.Remove(c, a.ID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
