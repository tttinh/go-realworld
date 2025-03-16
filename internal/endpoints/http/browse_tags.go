package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type batchTagRes struct {
	Tags []string `json:"tags"`
}

func (h *Handler) browseTags(c *gin.Context) {
	tags, err := h.articles.GetAllTags(c)
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, batchTagRes{Tags: tags})
}
