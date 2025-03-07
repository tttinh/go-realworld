package httpendpoints

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

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

func (h *Handler) loginUser(c *gin.Context) {
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
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}

func (h *Handler) getCurrentUser(c *gin.Context) {
	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}

func (h *Handler) updateCurrentUser(c *gin.Context) {
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
	res.User.Token, _ = h.token.generate(strconv.Itoa(u.ID))
	ok(c, res)
}

func (h *Handler) getProfile(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		error400(c, err)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	ok(c, res)
}

func (h *Handler) followUser(c *gin.Context) {
	followerID := 1
	followingUsername := c.Param("username")
	followingUser, err := h.users.GetProfile(c, followerID, followingUsername)
	if err != nil {
		error400(c, err)
		return
	}

	err = h.users.Follow(c, followerID, followingUser.ID)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res profileRes
	res.fromEntity(followingUser)
	res.Profile.Following = true
	ok(c, res)
}

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
