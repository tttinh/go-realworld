package httpendpoints

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func abort(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.AbortWithError(http.StatusInternalServerError, err)
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

type commentRes struct {
	Comment struct {
		ID        int       `json:"id"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Author    struct {
			Name      string `json:"username"`
			Bio       string `json:"bio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		} `json:"author"`
	} `json:"comment"`
}

func (res *commentRes) fromEntity(c domain.Comment) {
	res.Comment.ID = c.ID
	res.Comment.Body = c.Body
	res.Comment.CreatedAt = c.CreatedAt
	res.Comment.UpdatedAt = c.UpdatedAt
	res.Comment.Author.Name = c.Author.Name
	res.Comment.Author.Bio = c.Author.Bio
	res.Comment.Author.Image = c.Author.Image
	res.Comment.Author.Following = c.Author.Following
}

type userRes struct {
	User struct {
		Name  string `json:"username"`
		Email string `json:"email"`
		Bio   string `json:"bio"`
		Image string `json:"image"`
		Token string `json:"token"`
	} `json:"user"`
}

func (res *userRes) fromEntity(u domain.User) {
	res.User.Name = u.Name
	res.User.Email = u.Email
	res.User.Bio = u.Bio
	res.User.Image = u.Image
	res.User.Token = "hihi"
}

type profileRes struct {
	Profile struct {
		Name      string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

func (res *profileRes) fromEntity(p domain.Profile) {
	res.Profile.Name = p.Name
	res.Profile.Bio = p.Bio
	res.Profile.Image = p.Image
	res.Profile.Following = p.Following
}
