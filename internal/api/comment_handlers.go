package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type CommentHandler struct {
	articles repo.Articles
	comments repo.Comments
}

func (h *CommentHandler) mount(router *gin.Engine) {
	router.GET("/articles/:slug/comments", h.browse)
	router.POST("/articles/:slug/comments", h.add)
	router.DELETE("/articles/:slug/comments/:id", h.delete)
}

func (h *CommentHandler) browse(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(c, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, newErrorRes(err))
	}

	comments, err := h.comments.FindByArticleId(c, a.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res commentsRes
	res.fromEntity(comments)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) add(c *gin.Context) {
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	comment := entity.Comment{
		Body: req.Comment.Body,
	}
	slug := c.Param("slug")
	comment, err := h.comments.Insert(c, slug, comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res commentRes
	res.fromEntity(comment)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) delete(c *gin.Context) {}
