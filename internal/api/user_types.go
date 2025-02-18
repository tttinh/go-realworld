package api

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
	} `json:"user"`
}
