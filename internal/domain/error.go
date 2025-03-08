package domain

import "errors"

var (
	ErrDuplicateKey  = errors.New("duplicated key")
	ErrNotFound      = errors.New("not found entity")
	ErrArticleUpdate = errors.New("one of updating fields must be non-empty")
)
