package pgrepo

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func toDomainError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.ConstraintName {
		case "articles_slug_key":
			return domain.ErrDuplicateKey
		}
	}

	return nil
}
