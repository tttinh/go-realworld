package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres/gendb"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Articles struct {
	*gendb.Queries
}

func NewArticles(db *pgx.Conn) *Articles {
	return &Articles{
		Queries: gendb.New(db),
	}
}

func (r *Articles) GetDetail(ctx context.Context, viewerID int, slug string) (domain.ArticleDetail, error) {
	row, err := r.GetArticleDetail(ctx, gendb.GetArticleDetailParams{
		Slug:     slug,
		ViewerID: int64(viewerID),
	})
	if err != nil {
		return domain.ArticleDetail{}, err
	}

	return domain.ArticleDetail{
		Article: domain.Article{
			ID:          int(row.ID),
			Slug:        row.Slug,
			Title:       row.Title,
			Description: row.Description,
			Body:        row.Body,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
		},
		Favorited:      row.Favorited,
		FavoritesCount: int(row.FavoritesCount),
		Author: domain.Author{
			Name:      row.Username,
			Bio:       row.Bio.String,
			Image:     row.Image.String,
			Following: row.Following,
		},
	}, nil
}

func (r *Articles) GetBySlug(ctx context.Context, slug string) (domain.Article, error) {
	row, err := r.GetArticleBySlug(ctx, slug)
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:          int(row.ID),
		AuthorID:    int(row.AuthorID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

func (r *Articles) Add(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := gendb.InsertArticleParams{
		AuthorID:    int64(a.AuthorID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
	row, err := r.InsertArticle(ctx, param)
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:          int(row.ID),
		AuthorID:    int(row.AuthorID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

func (r *Articles) Edit(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := gendb.UpdateArticleParams{
		ID:          int64(a.ID),
		AuthorID:    int64(a.AuthorID),
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
		AuthorID:    int(row.AuthorID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

func (r *Articles) Remove(ctx context.Context, id int) error {
	return r.DeleteArticle(ctx, int64(id))
}

func (r *Articles) AddFavorite(ctx context.Context, userID int, articleID int) error {
	return nil
}

func (r *Articles) RemoveFavorite(ctx context.Context, userID int, articleID int) error {
	return nil
}
