package httpendpoints

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type registerUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (h *Handler) registerUser(c *gin.Context) {
	var req registerUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	u := domain.NewUser(req.User.Name, req.User.Email, req.User.Password)
	u, err := h.users.Add(c, u)
	if err != nil {
		error400(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}
