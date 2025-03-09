package httpendpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCurrentUser(c *gin.Context) {
	u, err := h.users.GetByID(c, 1)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.jwt.Generate(strconv.Itoa(u.ID))
	c.JSON(http.StatusOK, res)
}
