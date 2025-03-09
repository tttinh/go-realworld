package domain

import "errors"

var (
	ErrForbidden     = errors.New("access forbidden")
	ErrNotFound      = errors.New("not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrDuplicateKey  = errors.New("duplicated key")
	ErrArticleUpdate = errors.New("one of updating fields must be non-empty")
)
