package httpendpoints

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getProfile(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		error400(c, err)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	ok(c, res)
}
