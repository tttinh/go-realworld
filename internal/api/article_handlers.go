package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct{}

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
	body := struct {
		Article CreateArticleReq `json:"article"`
	}{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	req := body.Article
	res := ArticleRes{
		Slug:           "abc",
		Title:          req.Title,
		Description:    req.Description,
		Body:           req.Body,
		Tags:           req.Tags,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Favorited:      false,
		FavoritesCount: 0,
	}

	c.JSON(http.StatusOK, gin.H{"article": res})
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
