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

func NewArticle(authorID int, title, description, body string, tags []string) Article {
	return Article{
		AuthorID:    authorID,
		Slug:        createSlug(title),
		Title:       title,
		Description: description,
		Body:        body,
		Tags:        tags,
	}
}

func (a *Article) Update(title string, description string, body string) error {
	if len(title) == 0 && len(description) == 0 && len(body) == 0 {
		return ErrArticleUpdate
	}

	if a.Title != title && len(title) > 0 {
		a.Title = title
		a.Slug = createSlug(title)
	}

	if len(description) > 0 {
		a.Description = description
	}

	if len(body) > 0 {
		a.Body = body
	}

	return nil
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
