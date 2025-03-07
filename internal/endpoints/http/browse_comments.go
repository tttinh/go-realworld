package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseComments(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error404(c)
		return
	}

	comments, err := h.articles.GetAllComments(c, a.ID)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res commentsRes
	res.fromEntity(comments)
	ok(c, res)
}
