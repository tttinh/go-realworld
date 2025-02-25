package domain

import (
	"regexp"
	"strings"
	"time"

	"math/rand/v2"
)

type Article struct {
	ID             int
	Slug           string
	Title          string
	Description    string
	Body           string
	Tags           []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount int
	Author         User
}

func NewArticle(authorID int, title, description, body string) Article {
	a := Article{
		Title:       title,
		Description: description,
		Body:        body,
		Author:      User{ID: authorID},
	}

	a.genSlug()

	return a
}

func (a *Article) genSlug() {
	slug := createSlug(a.Title)
	randomString := createRandomString(12)
	a.Slug = slug + "-" + randomString
}

func createSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace non-alphanumeric characters with a hyphen
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove consecutive hyphens and trailing hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}

func createRandomString(length int) string {
	rng := rand.New(rand.NewPCG(0, uint64(time.Now().UnixNano())))
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rng.IntN(len(chars))]
	}
	return string(result)
}
