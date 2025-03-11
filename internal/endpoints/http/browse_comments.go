package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type batchCommentsRes struct {
	Comments []commentRes `json:"comments"`
}

func (res *batchCommentsRes) fromEntity(comments []domain.Comment) {
	for _, c := range comments {
		res.Comments = append(res.Comments, commentRes{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
}

func (h *Handler) browseComments(c *gin.Context) {
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		abort(c, err)
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
