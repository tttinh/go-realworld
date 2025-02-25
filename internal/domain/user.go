package domain

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Bio      string
	Image    string
}

func NewUser(name, email, password string) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
