package httpport

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type userHandler struct {
	users domain.UserRepo
}

func (h *userHandler) register(c *gin.Context) {
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
	ok(c, res)
}

func (h *userHandler) login(c *gin.Context) {
	var req loginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	u, err := h.users.GetByEmail(c, req.User.Email)
	if err != nil {
		error400(c, err)
		return
	}

	if u.Password != req.User.Password {
		error400(c, ErrWrongPassword)
		return
	}

	var res userRes
	res.fromEntity(u)
	ok(c, res)
}

func (h *userHandler) read(c *gin.Context) {
	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	ok(c, res)
}

func (h *userHandler) edit(c *gin.Context) {
	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	u.Name = req.User.Name
	u.Email = req.User.Email
	u.Password = req.User.Password
	u.Bio = req.User.Bio
	u.Image = req.User.Image

	u, err = h.users.Edit(c, u)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res userRes
	res.fromEntity(u)
	ok(c, res)
}
