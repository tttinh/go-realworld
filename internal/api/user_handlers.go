package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type userHandler struct {
	users repo.Users
}

func (h *userHandler) mount(router *gin.RouterGroup) {
	router.POST("/users/login", h.login)
	router.POST("/users", h.register)
	router.GET("/user", h.read)
	router.PUT("/user", h.edit)
}

func (h *userHandler) register(c *gin.Context) {
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

func (h *userHandler) login(c *gin.Context) {
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

func (h *userHandler) read(c *gin.Context) {
	u, err := h.users.FindById(c, 1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	var res userRes
	res.fromEntity(u)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) edit(c *gin.Context) {
	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	u, err := h.users.FindById(c, 1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorRes(err))
	}

	u.Name = req.User.Name
	u.Email = req.User.Email
	u.Password = req.User.Password
	u.Bio = req.User.Bio
	u.Image = req.User.Image

	u, err = h.users.Update(c, u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorRes(err))
	}

	var res userRes
	res.fromEntity(u)
	c.JSON(http.StatusOK, res)
}
