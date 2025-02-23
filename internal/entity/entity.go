package entity

import (
	"strings"
	"time"
)

type User struct {
	Id       int
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

type Article struct {
	Id          int
	Slug        string
	Title       string
	Description string
	Body        string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewArticle(title, description, body string) Article {
	a := Article{
		Title:       title,
		Description: description,
		Body:        body,
	}

	a.GenerateSlug()

	return a
}

func (a *Article) GenerateSlug() {
	a.Slug = strings.ToLower(strings.TrimSpace(a.Title))
}

type Comment struct {
	Id        int
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
