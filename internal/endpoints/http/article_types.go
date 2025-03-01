package httpport

import (
	"time"

	"github.com/tinhtt/go-realworld/internal/domain"
)

type createArticleReq struct {
	Article struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Body        string   `json:"body"`
		Tags        []string `json:"tagList"`
	} `json:"article"`
}

type updateArticleReq struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

type articleRes struct {
	Article struct {
		Slug           string    `json:"slug"`
		Title          string    `json:"title"`
		Description    string    `json:"description"`
		Body           string    `json:"body"`
		Tags           []string  `json:"tagList"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		Favorited      bool      `json:"favorited"`
		FavoritesCount int       `json:"favoritesCount"`
		Author         struct {
			Username  string `json:"username"`
			Bio       string `json:"bio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		} `json:"author"`
	} `json:"article"`
}

func (res *articleRes) fromEntity(a domain.ArticleDetail) {
	res.Article.Slug = a.Slug
	res.Article.Title = a.Title
	res.Article.Description = a.Description
	res.Article.Body = a.Body
	res.Article.Tags = a.Tags
	res.Article.CreatedAt = a.CreatedAt
	res.Article.UpdatedAt = a.UpdatedAt
	res.Article.Favorited = a.Favorited
	res.Article.FavoritesCount = a.FavoritesCount
	res.Article.Author.Username = a.Author.Name
	res.Article.Author.Bio = a.Author.Bio
	res.Article.Author.Image = a.Author.Image
	res.Article.Author.Following = a.Author.Following
}
