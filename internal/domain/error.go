package domain

import "errors"

var (
	ErrArticleUpdate = errors.New("one of updating fields must be non-empty")
)
