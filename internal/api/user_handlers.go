package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (h *UserHandler) Mount(router *gin.Engine) {
	router.POST("/users/login", h.Login)
	router.POST("/users", h.Register)
	router.GET("/users", h.Read)
	router.PUT("/users", h.Edit)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorRes(err))
	}

	c.JSON(http.StatusOK, "hihi")
}

func (h *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}

func (h *UserHandler) Read(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}

func (h *UserHandler) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}
