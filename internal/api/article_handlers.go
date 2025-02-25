package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type articleHandler struct {
	articles repo.Articles
}

func (h *articleHandler) mount(router *gin.RouterGroup) {
	router.GET("/articles/feed", h.browseFeed)
	router.GET("/articles", h.browse)
	router.GET("/articles/:slug", h.read)
	router.PUT("/articles/:slug", h.edit)
	router.POST("/articles", h.add)
	router.DELETE("/articles/:slug", h.delete)
}

func (h *articleHandler) browseFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *articleHandler) browse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *articleHandler) read(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(c, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) edit(c *gin.Context) {
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(c, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	a.Title = req.Article.Title
	a.Description = req.Article.Description
	a.Body = req.Article.Body

	a, err = h.articles.Update(c, a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) add(c *gin.Context) {
	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	a := entity.NewArticle(
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Insert(c, a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) delete(c *gin.Context) {
	slug := c.Param("slug")
	if err := h.articles.Delete(c, slug); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}
}
