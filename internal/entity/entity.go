package entity

import "time"

type User struct {
	Id        int
	Email     string
	Username  string
	Bio       string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	Id          int
	Slug        string
	Title       string
	Description string
	Body        string
	TagList     []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Comment struct {
	Id        int
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
