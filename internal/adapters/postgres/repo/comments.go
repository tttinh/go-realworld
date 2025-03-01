package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres/gendb"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Comments struct {
	*gendb.Queries
}

func NewComments(db *pgx.Conn) *Comments {
	return &Comments{
		Queries: gendb.New(db),
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
