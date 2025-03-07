package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) favoriteArticle(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if err := h.articles.AddFavorite(c, authorID, a.ID); err != nil {
		log.Println(err)
		error500(c)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, a.Slug)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	ok(c, res)
}
