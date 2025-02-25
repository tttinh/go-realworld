package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tinhtt/go-realworld/internal/domain"
	pgdb "github.com/tinhtt/go-realworld/internal/infra/postgres"
)

func articleFromDB(a pgdb.Article) domain.Article {
	return domain.Article{
		ID:          int(a.ID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		// Tags:    a.Image.String,
		CreatedAt:      a.CreatedAt.Time,
		UpdatedAt:      a.UpdatedAt.Time,
		Favorited:      false,
		FavoritesCount: 0,
		Author: domain.User{
			ID: 1,
		},
	}
}

type Articles struct {
	*pgdb.Queries
}

func NewArticles(db *pgx.Conn) *Articles {
	return &Articles{
		Queries: pgdb.New(db),
	}
}

func (r *Articles) FindBySlug(ctx context.Context, slug string) (domain.Article, error) {
	dbArticle, err := r.GetArticleBySlug(ctx, slug)
	if err != nil {
		return domain.Article{}, err
	}

	return articleFromDB(dbArticle), nil
}

func (r *Articles) Insert(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := pgdb.CreateArticleParams{
		AuthorID:    int64(a.Author.ID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
	dbArticle, err := r.CreateArticle(ctx, param)
	if err != nil {
		return domain.Article{}, err
	}

	return articleFromDB(dbArticle), nil
}

func (r *Articles) Update(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := pgdb.UpdateArticleParams{
		ID:          int64(a.ID),
		AuthorID:    int64(a.Author.ID),
		Slug:        pgtype.Text{String: a.Slug, Valid: len(a.Slug) > 0},
		Title:       pgtype.Text{String: a.Title, Valid: len(a.Title) > 0},
		Description: pgtype.Text{String: a.Description, Valid: len(a.Description) > 0},
		Body:        pgtype.Text{String: a.Body, Valid: len(a.Body) > 0},
	}
	dbArticle, err := r.UpdateArticle(ctx, param)
	if err != nil {
		return domain.Article{}, err
	}

	return articleFromDB(dbArticle), nil
}

func (r *Articles) Delete(ctx context.Context, slug string) error {
	return nil
}
