package httpendpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCurrentUser(c *gin.Context) {
	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}
