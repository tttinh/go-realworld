package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Articles struct {
	*sqlc.Queries
	db *pgx.Conn
}

func NewArticles(db *pgx.Conn) *Articles {
	return &Articles{
		Queries: sqlc.New(db),
		db:      db,
	}
}

func (r *Articles) GetDetail(ctx context.Context, viewerID int, slug string) (domain.ArticleDetail, error) {
	row, err := r.FetchArticleDetail(ctx, sqlc.FetchArticleDetailParams{
		Slug:     slug,
		ViewerID: int64(viewerID),
	})
	if err != nil {
		return domain.ArticleDetail{}, toDomainError(err)
	}

	return domain.ArticleDetail{
		Article: domain.Article{
			ID:          int(row.ID),
			Slug:        row.Slug,
			Title:       row.Title,
			Description: row.Description,
			Body:        row.Body,
			Tags:        row.Tags,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
			Author: domain.Author{
				ID:        int(row.AuthorID),
				Name:      row.Username,
				Bio:       row.Bio.String,
				Image:     row.Image.String,
				Following: row.Following,
			},
		},
		Favorited:      row.Favorited,
		FavoritesCount: int(row.FavoritesCount),
	}, nil
}

func (r *Articles) Get(ctx context.Context, slug string) (domain.Article, error) {
	row, err := r.FetchArticleBySlug(ctx, slug)
	if err != nil {
		return domain.Article{}, toDomainError(err)
	}

	return domain.Article{
		ID:          int(row.ID),
		Slug:        row.Slug,
		Title:       row.Title,
		Description: row.Description,
		Body:        row.Body,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		Author: domain.Author{
			ID: int(row.AuthorID),
		},
	}, nil
}

func (r *Articles) Edit(ctx context.Context, a domain.Article) (domain.Article, error) {
	param := sqlc.UpdateArticleParams{
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
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
		Author: domain.Author{
			ID: int(row.AuthorID),
		},
	}, nil
}

func (r *Articles) Remove(ctx context.Context, id int) error {
	return r.DeleteArticle(ctx, int64(id))
}

func (r *Articles) AddFavorite(ctx context.Context, userID int, articleID int) error {
	return r.InsertFavorite(ctx, sqlc.InsertFavoriteParams{
		UserID:    int64(userID),
		ArticleID: int64(articleID),
	})
}

func (r *Articles) RemoveFavorite(ctx context.Context, userID int, articleID int) error {
	return r.DeleteFavorite(ctx, sqlc.DeleteFavoriteParams{
		UserID:    int64(userID),
		ArticleID: int64(articleID),
	})
}

func (r *Articles) GetAllTags(ctx context.Context) ([]string, error) {
	return r.FetchAllTags(ctx)
}
