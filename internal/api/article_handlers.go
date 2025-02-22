package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type ArticleHandler struct {
	articles repo.Articles
}

func (h *ArticleHandler) Mount(router *gin.Engine) {
	router.GET("/articles/feed", h.GetFeed)
	router.GET("/articles", h.Browse)
	router.GET("/articles/:slug", h.Read)
	router.PUT("/articles/:slug", h.Edit)
	router.POST("/articles", h.Add)
	router.DELETE("/articles/:slug", h.Delete)
}

func (h *ArticleHandler) GetFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) Browse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) Add(c *gin.Context) {
	var req CreateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	a := entity.NewArticle(
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Insert(a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) Read(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) Edit(c *gin.Context) {
	var req UpdateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	a.Title = req.Article.Title
	a.Description = req.Article.Description
	a.Body = req.Article.Body

	a, err = h.articles.Update(a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	slug := c.Param("slug")
	if err := h.articles.Delete(slug); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}
}
