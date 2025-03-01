package httpport

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type commentHandler struct {
	articles domain.ArticleRepo
	comments domain.CommentRepo
}

func (h *commentHandler) browse(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		error404(c)
		return
	}

	comments, err := h.comments.FindAllByArticleId(c, a.ID)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res commentsRes
	res.fromEntity(comments)
	ok(c, res)
}

func (h *commentHandler) add(c *gin.Context) {
	authorID := 1
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		error404(c)
		return
	}

	comment := domain.Comment{
		AuthorID:  authorID,
		ArticleID: a.ID,
		Body:      req.Comment.Body,
	}
	comment, err = h.comments.Insert(c, comment)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res commentRes
	res.fromEntity(comment)
	ok(c, res)
}

func (h *commentHandler) delete(c *gin.Context) {}
