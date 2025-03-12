package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleQueries struct {
	Tag       *string `form:"tag"`
	Author    *string `form:"author"`
	Favorited *string `form:"favorited"`
	Offset    int     `form:"offset,default=1"`
	Limit     int     `form:"limit,default=10"`
}

type batchArticleRes struct {
	Articles     []articleData `json:"articles"`
	ArticleCount int           `json:"articleCount"`
}

func (h *Handler) browseArticles(c *gin.Context) {
	var q articleQueries
	if err := c.ShouldBindQuery(&q); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, _ := h.jwt.GetUserID(c)
	if q.Author != nil {
		h.browseArticlesByAuthor(c, userID, q)
		return
	}

	if q.Favorited != nil {
		h.browseArticlesByFavorited(c, userID, q)
		return
	}

	if q.Tag != nil {
		h.browseArticlesByTag(c, userID, q)
		return
	}

	// articles, err := h.articles.GetAllArticles(c, userID, q.Offset, q.Limit)
	var res batchArticleRes
	c.JSON(http.StatusOK, res)
}

func (h *Handler) browseArticlesByTag(c *gin.Context, userID int, q articleQueries) {
	panic("unimplemented")
}

func (h *Handler) browseArticlesByFavorited(c *gin.Context, userID int, q articleQueries) {
	panic("unimplemented")
}

func (h *Handler) browseArticlesByAuthor(c *gin.Context, userID int, q articleQueries) {
	panic("unimplemented")
}
