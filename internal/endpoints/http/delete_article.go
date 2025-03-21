package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) deleteArticle(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		abortWithError(c, err)
		return
	}

	if userID != a.Author.ID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	if err := h.articles.Remove(c, a.ID); err != nil {
		abortWithError(c, err)
		return
	}
}
