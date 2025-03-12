package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type articleQueries struct {
	Tag       *string `form:"tag"`
	Author    *string `form:"author"`
	Favorited *string `form:"favorited"`
	Offset    int     `form:"offset,default=0"`
	Limit     int     `form:"limit,default=10"`
}

type batchArticleRes struct {
	Articles      []articleData `json:"articles"`
	ArticlesCount int           `json:"articlesCount"`
}

func (res *batchArticleRes) fromEntity(items []domain.ArticleDetail, total int) {
	res.ArticlesCount = total
	res.Articles = []articleData{}
	for i := range items {
		var data articleData
		data.fromEntity(items[i])
		res.Articles = append(res.Articles, data)
	}
}

func (h *Handler) browseArticles(c *gin.Context) {
	var q articleQueries
	if err := c.ShouldBindQuery(&q); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, _ := h.jwt.GetUserID(c)
	// if q.Author != nil {
	// 	h.browseArticlesByAuthor(c, userID, q)
	// 	return
	// }

	// if q.Favorited != nil {
	// 	h.browseArticlesByFavorited(c, userID, q)
	// 	return
	// }

	// if q.Tag != nil {
	// 	h.browseArticlesByTag(c, userID, q)
	// 	return
	// }

	articles, err := h.articles.GetAllArticles(c, userID, q.Offset, q.Limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res batchArticleRes
	res.fromEntity(articles, 10)
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
