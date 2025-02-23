package api

import "github.com/tinhtt/go-realworld/internal/entity"

type RegisterUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type LoginUserReq struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type UpdateUserReq struct {
	User struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

type UserRes struct {
	User struct {
		Name  string `json:"username"`
		Email string `json:"email"`
		Bio   string `json:"bio"`
		Image string `json:"image"`
		Token string `json:"token"`
	} `json:"user"`
}

func NewUserRes(u entity.User) UserRes {
	var res UserRes
	res.User.Name = u.Name
	res.User.Email = u.Email
	res.User.Bio = u.Bio
	res.User.Image = u.Image
	res.User.Token = "hihi"

	return res
}
