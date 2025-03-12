package domain

type Profile struct {
	ID        int
	Name      string
	Bio       string
	Image     string
	Following bool
}

type Author struct {
	ID        int
	Name      string
	Bio       string
	Image     string
	Following bool
}

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
