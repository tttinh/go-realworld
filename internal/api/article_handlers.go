package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type ArticleHandler struct {
	articles repo.Article
}

func (h *ArticleHandler) Mount(router *gin.Engine) {
	router.GET("/articles/feed", h.GetFeed)
	router.GET("/articles", h.ListArticles)
	router.POST("/articles", h.CreateArticle)
	router.GET("/articles/:slug", h.GetArticle)
	router.PUT("/articles/:slug", h.UpdateArticle)
	router.DELETE("/articles/:slug", h.DeleteArticle)
}

func (h *ArticleHandler) GetFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) ListArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var req CreateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := entity.NewArticle(
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Insert(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	var req UpdateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.FindBySlug(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.Title = req.Article.Title
	a.Description = req.Article.Description
	a.Body = req.Article.Body

	a, err = h.articles.Update(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res ArticleRes
	res.FromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	slug := c.Param("slug")
	if err := h.articles.Delete(slug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
