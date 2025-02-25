package repo

import (
	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/domain"
	pgrepo "github.com/tinhtt/go-realworld/internal/repo/postgres"
)

func NewPostgresUsers(db *pgx.Conn) domain.UserRepo {
	return pgrepo.NewUsers(db)
}

func NewPostgresArticles(db *pgx.Conn) domain.ArticleRepo {
	return pgrepo.NewArticles(db)
}

func NewPostgresComments(db *pgx.Conn) domain.CommentRepo {
	return pgrepo.NewComments(db)
}
