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
	comment, err := h.articles.GetComment(c, commentID)
	if err != nil {
		abort(c, err)
	}

	if userID != comment.AuthorID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	if err := h.articles.RemoveComment(c, commentID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
