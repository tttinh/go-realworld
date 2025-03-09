package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) deleteArticle(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		abort(c, err)
		return
	}

	if authorID != a.AuthorID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	if err := h.articles.Remove(c, a.ID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
