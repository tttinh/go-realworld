package domain

import (
	"math/rand/v2"
	"regexp"
	"strings"
	"time"
)

func makeSlugWithRandomString(s string) string {
	return makeSlug(s) + "-" + makeRandomString(12)
}

func makeSlug(s string) string {
	// Convert to lowercase
	slug := strings.ToLower(s)

	// Replace non-alphanumeric characters with a hyphen
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove consecutive hyphens and trailing hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}

func makeRandomString(length int) string {
	rng := rand.New(rand.NewPCG(0, uint64(time.Now().UnixNano())))
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rng.IntN(len(chars))]
	}
	return string(result)
}
