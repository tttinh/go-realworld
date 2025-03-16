package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type feedQuery struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit,default=10"`
}

func (h *Handler) browseFeed(c *gin.Context) {
	var q feedQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, _ := h.jwt.GetUserID(c)
	articles, err := h.articles.GetFeed(c, userID, q.Offset, q.Limit)
	if err != nil {
		abortWithError(c, err)
		return
	}

	var res batchArticleRes
	res.fromEntity(articles)
	c.JSON(http.StatusOK, res)
}
