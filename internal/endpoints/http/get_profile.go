package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getProfile(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		abort(c, err)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	c.JSON(http.StatusOK, res)
}
