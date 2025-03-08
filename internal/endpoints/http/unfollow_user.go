package httpendpoints

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) unfollowUser(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		error400(c, err)
		return
	}

	err = h.users.Unfollow(c, followerID, followingUser.ID)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	res.Profile.Following = false
	ok(c, res)
}
