package httpendpoints

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateCurrentUser(c *gin.Context) {
	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	u.Name = req.User.Name
	u.Email = req.User.Email
	u.Password = req.User.Password
	u.Bio = req.User.Bio
	u.Image = req.User.Image

	u, err = h.users.Edit(c, u)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}
