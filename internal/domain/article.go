package domain

import (
	"time"
)

type Article struct {
	Author      Author
	ID          int
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
		Slug:        makeSlug(title),
		Title:       title,
		Description: description,
		Body:        body,
		Tags:        tags,
		Author:      Author{ID: authorID},
	}
}

func (a *Article) Update(title string, description string, body string) error {
	if len(title) == 0 && len(description) == 0 && len(body) == 0 {
		return ErrArticleUpdate
	}

	if a.Title != title && len(title) > 0 {
		a.Title = title
		a.Slug = makeSlug(title)
	}

	if len(description) > 0 {
		a.Description = description
	}

	if len(body) > 0 {
		a.Body = body
	}

	return nil
}

func (a *Article) NewSlug() {
	a.Slug = makeSlugWithRandomString(a.Title)
}

type ArticleDetail struct {
	Article
	Favorited      bool
	FavoritesCount int
}

type ArticleList struct {
	Articles []ArticleDetail
	Total    int
}
