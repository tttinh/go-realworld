package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/domain"
	pgdb "github.com/tinhtt/go-realworld/internal/infra/postgres"
)

type Comments struct {
	*pgdb.Queries
}

func NewComments(db *pgx.Conn) *Comments {
	return &Comments{
		Queries: pgdb.New(db),
	}
}

func (repo *Comments) FindByArticleId(ctx context.Context, id int) ([]domain.Comment, error) {
	return nil, nil
}

func (repo *Comments) Insert(ctx context.Context, slug string, c domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}

func (repo *Comments) Delete(ctx context.Context, id int) error {
	return nil
}
