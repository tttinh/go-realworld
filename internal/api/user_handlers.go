package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type UserHandler struct {
	users repo.Users
}

func (h *UserHandler) mount(router *gin.Engine) {
	router.POST("/users/login", h.login)
	router.POST("/users", h.register)
	router.GET("/users", h.read)
	router.PUT("/users", h.edit)
}

func (h *UserHandler) register(c *gin.Context) {
	var req registerUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	u := entity.NewUser(req.User.Name, req.User.Email, req.User.Password)
	u, err := h.users.Insert(c, u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorRes(err))
	}

	var res userRes
	res.fromEntity(u)
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) login(c *gin.Context) {
	var req loginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	u, err := h.users.FindByEmail(c, req.User.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	if u.Password != req.User.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(errors.New("wrong password")))
	}

	var res userRes
	res.fromEntity(u)
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) read(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}

func (h *UserHandler) edit(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}
