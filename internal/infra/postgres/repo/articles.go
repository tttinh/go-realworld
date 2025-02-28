package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tinhtt/go-realworld/internal/domain"
	"github.com/tinhtt/go-realworld/internal/infra/postgres/gendb"
)

type Articles struct {
	*gendb.Queries
}

func NewArticles(db *pgx.Conn) *Articles {
	return &Articles{
		Queries: gendb.New(db),
	}
}

func (r *Articles) FindBySlug(ctx context.Context, viewerID int, slug string) (domain.Article, error) {
	row, err := r.GetArticleBySlug(ctx, gendb.GetArticleBySlugParams{
		Slug:     slug,
		ViewerID: int64(viewerID),
	})
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:          int(row.ID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		// Tags:    a.Image.String,
		CreatedAt:      row.CreatedAt.Time,
		UpdatedAt:      row.UpdatedAt.Time,
		Favorited:      row.Favorited,
		FavoritesCount: int(row.FavoritesCount),
		Author: domain.Author{
			ID:        int(row.AuthorID),
			Bio:       row.Bio.String,
			Image:     row.Image.String,
			Following: row.Following,
		},
	}, nil
}

func (r *Articles) Insert(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := gendb.CreateArticleParams{
		AuthorID:    int64(a.Author.ID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
	row, err := r.CreateArticle(ctx, param)
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:          int(row.ID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		// Tags:    a.Image.String,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		Author: domain.Author{
			ID: int(row.AuthorID),
		},
	}, nil
}

func (r *Articles) Update(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := gendb.UpdateArticleParams{
		ID:          int64(a.ID),
		AuthorID:    int64(a.Author.ID),
		Slug:        pgtype.Text{String: a.Slug, Valid: len(a.Slug) > 0},
		Title:       pgtype.Text{String: a.Title, Valid: len(a.Title) > 0},
		Description: pgtype.Text{String: a.Description, Valid: len(a.Description) > 0},
		Body:        pgtype.Text{String: a.Body, Valid: len(a.Body) > 0},
	}
	row, err := r.UpdateArticle(ctx, param)
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:          int(row.ID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		// Tags:    a.Image.String,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		Author: domain.Author{
			ID: int(row.AuthorID),
		},
	}, nil
}

func (r *Articles) Delete(ctx context.Context, slug string) error {
	return nil
}
