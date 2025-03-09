package httpendpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type loginUserReq struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (h *Handler) loginUser(c *gin.Context) {
	var req loginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, err := h.users.GetByEmail(c, req.User.Email)
	if err != nil {
		abort(c, err)
		return
	}

	if u.Password != req.User.Password {
		c.AbortWithError(http.StatusBadRequest, domain.ErrWrongPassword)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.jwt.Generate(strconv.Itoa(u.ID))
	c.JSON(http.StatusOK, res)
}
