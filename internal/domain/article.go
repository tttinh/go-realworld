package domain

import (
	"time"
)

type Article struct {
	ID          int
	AuthorID    int
	Slug        string
	Title       string
	Description string
	Body        string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewArticle(authorID int, title, description, body string) Article {
	return Article{
		AuthorID:    authorID,
		Slug:        createSlug(title),
		Title:       title,
		Description: description,
		Body:        body,
	}
}

type Author struct {
	Name      string
	Bio       string
	Image     string
	Following bool
}

type ArticleDetail struct {
	Article
	Favorited      bool
	FavoritesCount int
	Author         Author
}
