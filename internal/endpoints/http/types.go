package httpendpoints

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
	"github.com/tinhtt/go-realworld/internal/pkg"
)

func abortWithError(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.AbortWithError(http.StatusInternalServerError, pkg.NewError(err))
}

type userRes struct {
	User struct {
		Name  string `json:"username"`
		Email string `json:"email"`
		Bio   string `json:"bio"`
		Image string `json:"image"`
		Token string `json:"token"`
	} `json:"user"`
}

func (res *userRes) fromEntity(u domain.User) {
	res.User.Name = u.Name
	res.User.Email = u.Email
	res.User.Bio = u.Bio
	res.User.Image = u.Image
	res.User.Token = "hihi"
}

type profileRes struct {
	Profile struct {
		Name      string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

func (res *profileRes) fromEntity(p domain.Profile) {
	res.Profile.Name = p.Name
	res.Profile.Bio = p.Bio
	res.Profile.Image = p.Image
	res.Profile.Following = p.Following
}
