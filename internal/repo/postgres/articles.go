package pgrepo

import (
	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/entity"
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

func (repo *Articles) FindBySlug(slug string) (entity.Article, error) {
	return entity.Article{}, nil
}

func (repo *Articles) Insert(a entity.Article) (entity.Article, error) {
	return entity.Article{}, nil
}

func (repo *Articles) Update(a entity.Article) (entity.Article, error) {
	return entity.Article{}, nil
}

func (repo *Articles) Delete(slug string) error {
	return nil
}
