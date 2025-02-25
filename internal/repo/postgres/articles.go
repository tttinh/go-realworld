package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/domain"
	pgdb "github.com/tinhtt/go-realworld/internal/infra/postgres"
)

type Articles struct {
	*pgdb.Queries
}

func NewArticles(db *pgx.Conn) *Articles {
	return &Articles{
		Queries: pgdb.New(db),
	}
}

func (repo *Articles) FindBySlug(ctx context.Context, slug string) (domain.Article, error) {
	return domain.Article{}, nil
}

func (repo *Articles) Insert(ctx context.Context, a domain.Article) (domain.Article, error) {
	return domain.Article{}, nil
}

func (repo *Articles) Update(ctx context.Context, a domain.Article) (domain.Article, error) {
	return domain.Article{}, nil
}

func (repo *Articles) Delete(ctx context.Context, slug string) error {
	return nil
}
