package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type batchCommentRes struct {
	Comments []commentData `json:"comments"`
}

func (res *batchCommentRes) fromEntity(comments []domain.Comment) {
	for _, c := range comments {
		var item commentData
		item.fromEntity(c)
		res.Comments = append(res.Comments, item)
	}
}

func (h *Handler) browseComments(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		abortWithError(c, err)
		return
	}

	comments, err := h.articles.GetAllComments(c, userID, a.ID)
	if err != nil {
		abortWithError(c, err)
		return
	}

	var res batchCommentRes
	res.fromEntity(comments)
	c.JSON(http.StatusOK, res)
}
