package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type createCommentReq struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func (h *Handler) createComment(c *gin.Context) {
	authorID := 1
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, ErrNotFound)
		return
	}

	comment := domain.Comment{
		AuthorID:  authorID,
		ArticleID: a.ID,
		Body:      req.Comment.Body,
	}
	comment, err = h.articles.AddComment(c, comment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res commentRes
	res.fromEntity(comment)
	c.JSON(http.StatusOK, res)
}
