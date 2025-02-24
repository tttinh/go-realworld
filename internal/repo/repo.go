package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/entity"
	pgrepo "github.com/tinhtt/go-realworld/internal/repo/postgres"
)

type Users interface {
	FindById(ctx context.Context, id int) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	Insert(ctx context.Context, u entity.User) (entity.User, error)
	Update(ctx context.Context, u entity.User) (entity.User, error)
}

type Articles interface {
	FindBySlug(ctx context.Context, slug string) (entity.Article, error)
	Insert(ctx context.Context, a entity.Article) (entity.Article, error)
	Update(ctx context.Context, a entity.Article) (entity.Article, error)
	Delete(ctx context.Context, slug string) error
}

type Comments interface {
	FindByArticleId(ctx context.Context, id int) ([]entity.Comment, error)
	Insert(ctx context.Context, slug string, c entity.Comment) (entity.Comment, error)
	Delete(ctx context.Context, id int) error
}

func NewPostgresUsers(db *pgx.Conn) Users {
	return pgrepo.NewUsers(db)
}

func NewPostgresArticles(db *pgx.Conn) Articles {
	return pgrepo.NewArticles(db)
}

func NewPostgresComments(db *pgx.Conn) Comments {
	return pgrepo.NewComments(db)
}
