package httpendpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) loginUser(c *gin.Context) {
	var req loginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	u, err := h.users.GetByEmail(c, req.User.Email)
	if err != nil {
		error400(c, err)
		return
	}

	if u.Password != req.User.Password {
		error400(c, ErrWrongPassword)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}
