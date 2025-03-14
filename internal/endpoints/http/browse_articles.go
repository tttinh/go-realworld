package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type articleQuery struct {
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

func (res *batchArticleRes) fromEntity(al domain.ArticleList) {
	res.ArticlesCount = al.Total
	res.Articles = []articleData{}
	for _, a := range al.Articles {
		var data articleData
		data.fromEntity(a)
		res.Articles = append(res.Articles, data)
	}
}

func (h *Handler) browseArticles(c *gin.Context) {
	var q articleQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, _ := h.jwt.GetUserID(c)
	var articles domain.ArticleList
	var err error
	switch {
	case q.Author != nil:
		articles, err = h.articles.GetAllArticlesByAuthor(c, userID, q.Offset, q.Limit, *q.Author)
	case q.Favorited != nil:
		articles, err = h.articles.GetAllArticlesByFavorited(c, userID, q.Offset, q.Limit, *q.Favorited)
	case q.Tag != nil:
		articles, err = h.articles.GetAllArticlesByTag(c, userID, q.Offset, q.Limit, *q.Tag)
	default:
		articles, err = h.articles.GetAllArticles(c, userID, q.Offset, q.Limit)
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res batchArticleRes
	res.fromEntity(articles)
	c.JSON(http.StatusOK, res)
}
