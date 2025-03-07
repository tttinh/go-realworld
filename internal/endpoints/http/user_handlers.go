package httpendpoints

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type userHandler struct {
	tokenSecret   string
	tokenDuration time.Duration
	users         domain.UserRepo
}

func (h *userHandler) createToken(id int) string {
	t, _ := generateToken(strconv.Itoa(id), h.tokenSecret, h.tokenDuration)
	return t
}
func (h *userHandler) registerUser(c *gin.Context) {
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
	res.User.Token = h.createToken(u.ID)
	ok(c, res)
}

func (h *userHandler) loginUser(c *gin.Context) {
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
	res.User.Token = h.createToken(u.ID)
	ok(c, res)
}

func (h *userHandler) getCurrentUser(c *gin.Context) {
	u, err := h.users.GetByID(c, 1)
	if err != nil {
		error400(c, err)
		return
	}

	var res userRes
	res.fromEntity(u)
	res.User.Token = h.createToken(u.ID)
	ok(c, res)
}

func (h *userHandler) updateCurrentUser(c *gin.Context) {
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
	res.User.Token = h.createToken(u.ID)
	ok(c, res)
}

func (h *userHandler) getProfile(c *gin.Context) {
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

func (h *userHandler) followUser(c *gin.Context) {
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

func (h *userHandler) unfollowUser(c *gin.Context) {
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
