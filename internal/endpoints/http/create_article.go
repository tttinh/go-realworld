package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) createArticle(c *gin.Context) {
	authorID := 1
	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	a := domain.NewArticle(
		authorID,
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Add(c, a)
	if err != nil {
		error400(c, err)
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
