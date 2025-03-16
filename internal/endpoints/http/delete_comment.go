package httpendpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (h *Handler) deleteComment(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	commentID, _ := strconv.Atoi(c.Param("id"))
	comment, err := h.articles.GetComment(c, userID, commentID)
	if err != nil {
		abortWithError(c, err)
	}

	if userID != comment.Author.ID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	if err := h.articles.RemoveComment(c, commentID); err != nil {
		abortWithError(c, err)
		return
	}
}
