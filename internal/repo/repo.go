package repo

import (
	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/entity"
	pgrepo "github.com/tinhtt/go-realworld/internal/repo/postgres"
)

type Users interface {
	FindById(id int) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Insert(u entity.User) (entity.User, error)
	Update(u entity.User) (entity.User, error)
}

type Articles interface {
	FindBySlug(slug string) (entity.Article, error)
	Insert(a entity.Article) (entity.Article, error)
	Update(a entity.Article) (entity.Article, error)
	Delete(slug string) error
}

type Comments interface {
	FindByArticleId(id int) ([]entity.Comment, error)
	Insert(slug string, c entity.Comment) (entity.Comment, error)
	Delete(id int) error
}

func NewPostgresArticles(db *pgx.Conn) Articles {
	return pgrepo.NewArticles(db)
}

func NewPostgresComments(db *pgx.Conn) Comments {
	return pgrepo.NewComments(db)
}
