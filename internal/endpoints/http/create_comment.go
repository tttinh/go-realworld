package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) createComment(c *gin.Context) {
	authorID := 1
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error404(c)
		return
	}

	comment := domain.Comment{
		AuthorID:  authorID,
		ArticleID: a.ID,
		Body:      req.Comment.Body,
	}
	comment, err = h.articles.AddComment(c, comment)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res commentRes
	res.fromEntity(comment)
	ok(c, res)
}
