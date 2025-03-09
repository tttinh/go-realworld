package httpendpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCurrentUser(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	u, err := h.users.GetByID(c, userID)
	if err != nil {
		abort(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.jwt.Generate(strconv.Itoa(u.ID))
	c.JSON(http.StatusOK, res)
}
