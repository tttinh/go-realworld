package pgrepo

import (
	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/entity"
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

func (repo *Comments) FindByArticleId(id int) ([]entity.Comment, error) {
	return nil, nil
}

func (repo *Comments) Insert(slug string, c entity.Comment) (entity.Comment, error) {
	return entity.Comment{}, nil
}

func (repo *Comments) Delete(id int) error {
	return nil
}
