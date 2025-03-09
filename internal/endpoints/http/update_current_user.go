package httpendpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

func (h *Handler) updateCurrentUser(c *gin.Context) {
	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, err := h.users.GetByID(c, 1)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u.Name = req.User.Name
	u.Email = req.User.Email
	u.Password = req.User.Password
	u.Bio = req.User.Bio
	u.Image = req.User.Image

	u, err = h.users.Edit(c, u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.jwt.Generate(strconv.Itoa(u.ID))
	c.JSON(http.StatusOK, res)
}
