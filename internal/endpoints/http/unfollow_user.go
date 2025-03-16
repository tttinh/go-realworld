package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) unfollowUser(c *gin.Context) {
	followerID, _ := h.jwt.GetUserID(c)
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		abortWithError(c, err)
		return
	}

	err = h.users.Unfollow(c, followerID, followingUser.ID)
	if err != nil {
		abortWithError(c, err)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	res.Profile.Following = false
	c.JSON(http.StatusOK, res)
}
