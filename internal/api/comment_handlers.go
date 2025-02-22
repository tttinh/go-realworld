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

func (h *CommentHandler) Mount(router *gin.Engine) {
	router.GET("/articles/:slug/comments", h.Browse)
	router.POST("/articles/:slug/comments", h.Add)
	router.DELETE("/articles/:slug/comments/:id", h.Delete)
}

func (h *CommentHandler) Browse(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, NewErrorRes(err))
	}

	comments, err := h.comments.FindByArticleId(a.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	res := commentsFromEntity(comments)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) Add(c *gin.Context) {
	var req CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	comment := entity.Comment{
		Body: req.Comment.Body,
	}
	slug := c.Param("slug")
	comment, err := h.comments.Insert(slug, comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	res := commentFromEntity(comment)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) Delete(c *gin.Context) {}
