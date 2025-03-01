package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type commentHandler struct {
	articles domain.ArticleRepo
	comments domain.CommentRepo
}

func (h *commentHandler) mount(router *gin.RouterGroup) {
	router.GET("/articles/:slug/comments", h.browse)
	router.POST("/articles/:slug/comments", h.add)
	router.DELETE("/articles/:slug/comments/:id", h.delete)
}

func (h *commentHandler) browse(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
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

func (h *commentHandler) add(c *gin.Context) {
	var req createCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	comment := domain.Comment{
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

func (h *commentHandler) delete(c *gin.Context) {}
