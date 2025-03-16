package httpendpoints

import (
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type articleData struct {
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
}

func (res *articleData) fromEntity(a domain.ArticleDetail) {
	res.Slug = a.Slug
	res.Title = a.Title
	res.Description = a.Description
	res.Body = a.Body
	res.CreatedAt = a.CreatedAt
	res.UpdatedAt = a.UpdatedAt
	res.Favorited = a.Favorited
	res.FavoritesCount = a.FavoritesCount
	res.Author.Username = a.Author.Name
	res.Author.Bio = a.Author.Bio
	res.Author.Image = a.Author.Image
	res.Author.Following = a.Author.Following

	res.Tags = []string{}
	if len(a.Tags) > 0 {
		slices.Sort(a.Tags)
		res.Tags = a.Tags
	}
}

type articleRes struct {
	Article articleData `json:"article"`
}

func (res *articleRes) fromEntity(a domain.ArticleDetail) {
	res.Article.fromEntity(a)
}

func (h *Handler) getArticle(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, userID, slug)
	if err != nil {
		abortWithError(c, err)
		return
	}

	var res articleRes
	res.fromEntity(a)
	c.JSON(http.StatusOK, res)
}
