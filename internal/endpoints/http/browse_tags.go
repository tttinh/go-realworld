package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type batchTagsRes struct {
	Tags []string `json:"tags"`
}

func (h *Handler) browseTags(c *gin.Context) {
	tags, err := h.articles.GetAllTags(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, batchTagsRes{Tags: tags})
}
