package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) browseFeed(c *gin.Context) {
	var res batchArticleRes
	res.fromEntity([]domain.ArticleDetail{}, 10)
	c.JSON(http.StatusOK, res)
}
