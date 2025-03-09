package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) browseComments(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, ErrNotFound)
		return
	}

	comments, err := h.articles.GetAllComments(c, a.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res batchCommentsRes
	res.fromEntity(comments)
	c.JSON(http.StatusOK, res)
}
