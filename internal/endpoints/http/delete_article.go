package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteArticle(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if authorID != a.AuthorID {
		error403(c)
		return
	}

	if err := h.articles.Remove(c, a.ID); err != nil {
		log.Println(err)
		error500(c)
		return
	}
}
