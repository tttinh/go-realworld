package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/entity"
	"github.com/tinhtt/go-realworld/internal/repo"
)

type UserHandler struct {
	users repo.Users
}

func (h *UserHandler) Mount(router *gin.Engine) {
	router.POST("/users/login", h.Login)
	router.POST("/users", h.Register)
	router.GET("/users", h.Read)
	router.PUT("/users", h.Edit)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err.Error()))
	}

	u := entity.NewUser(req.User.Name, req.User.Email, req.User.Password)
	u, err := h.users.Insert(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorRes(err.Error()))
	}

	c.JSON(http.StatusOK, NewUserRes(u))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err.Error()))
	}

	u, err := h.users.FindByEmail(req.User.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err.Error()))
	}

	if u.Password != req.User.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes("wrong password"))
	}

	c.JSON(http.StatusOK, NewUserRes(u))
}

func (h *UserHandler) Read(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}

func (h *UserHandler) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}
