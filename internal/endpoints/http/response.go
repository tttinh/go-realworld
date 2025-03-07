package httpendpoints

import (
	"time"

	"github.com/tinhtt/go-realworld/internal/domain"
)

type createCommentReq struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

type commentRes struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"author"`
}

func (res *commentRes) fromEntity(c domain.Comment) {
	res.ID = c.ID
	res.Body = c.Body
	res.CreatedAt = c.CreatedAt
	res.UpdatedAt = c.UpdatedAt
}

type commentsRes struct {
	Comments []commentRes `json:"comments"`
}

func (res *commentsRes) fromEntity(comments []domain.Comment) {
	for _, c := range comments {
		res.Comments = append(res.Comments, commentRes{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
}

type registerUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type loginUserReq struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type updateUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
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
