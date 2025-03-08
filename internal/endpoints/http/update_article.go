package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
)

type updateArticleReq struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (h *Handler) updateArticle(c *gin.Context) {
	authorID := 1
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

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

	err = a.Update(req.Article.Title, req.Article.Description, req.Article.Body)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	a, err = h.articles.Edit(c, a)
	if err != nil {
		error400(c, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, slug)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	ok(c, res)
}
