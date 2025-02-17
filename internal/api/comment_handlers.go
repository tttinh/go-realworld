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
	router.GET("/articles/:slug/comments", h.GetAll)
	router.POST("/articles/:slug/comments", h.Create)
	router.DELETE("/articles/:slug/comments/:id", h.Delete)
}

func (h *CommentHandler) GetAll(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	comments, err := h.comments.FindByArticleId(a.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := commentsFromEntity(comments)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	comment := entity.Comment{
		Body: req.Comment.Body,
	}
	slug := c.Param("slug")
	comment, err := h.comments.Insert(slug, comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := commentFromEntity(comment)
	c.JSON(http.StatusOK, res)
}

func (h *CommentHandler) Delete(c *gin.Context) {}
