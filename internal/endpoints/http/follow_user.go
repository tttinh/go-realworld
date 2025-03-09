package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) followUser(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		abort(c, err)
		return
	}

	err = h.users.Follow(c, followerID, followingUser.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	res.Profile.Following = true
	c.JSON(http.StatusOK, res)
}
