package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type articleHandler struct {
	articles domain.ArticleRepo
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
	viewerID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, viewerID, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(a)

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) edit(c *gin.Context) {
	authorID := 1
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	if authorID != a.AuthorID {
		c.AbortWithStatusJSON(http.StatusForbidden, newErrorRes(ErrAccessForbidden))
	}

	a.Title = req.Article.Title
	a.Description = req.Article.Description
	a.Body = req.Article.Body
	a, err = h.articles.Update(c, a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	detail, err := h.articles.GetDetail(c, authorID, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) add(c *gin.Context) {
	authorID := 1
	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	a := domain.NewArticle(
		authorID,
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Insert(c, a)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	detail, err := h.articles.GetDetail(c, authorID, a.Slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) delete(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	if authorID != a.AuthorID {
		c.AbortWithStatusJSON(http.StatusForbidden, newErrorRes(ErrAccessForbidden))
	}

	if err := h.articles.Delete(c, slug); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}
}
