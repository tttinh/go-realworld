package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) favoriteArticle(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		abortWithError(c, err)
		return
	}

	if err := h.articles.AddFavorite(c, userID, a.ID); err != nil {
		abortWithError(c, err)
		return
	}

	detail, err := h.articles.GetDetail(c, userID, a.Slug)
	if err != nil {
		abortWithError(c, err)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}
