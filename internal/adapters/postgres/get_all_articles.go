package postgres

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetAllArticles(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
) ([]domain.ArticleDetail, error) {
	results := []domain.ArticleDetail{}
	rows, err := r.FetchAllArticles(ctx, sqlc.FetchAllArticlesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
		UserID: int64(viewerID),
	})
	if err != nil {
		return results, err
	}

	for _, r := range rows {
		results = append(results, domain.ArticleDetail{
			Article: domain.Article{
				ID:          int(r.ID),
				Slug:        r.Slug,
				Title:       r.Title,
				Description: r.Description,
				Body:        r.Body,
				Tags:        r.Tags,
				CreatedAt:   r.CreatedAt.Time,
				UpdatedAt:   r.UpdatedAt.Time,
				Author: domain.Author{
					ID:        int(r.AuthorID),
					Name:      r.AuthorName,
					Bio:       r.AuthorBio.String,
					Image:     r.AuthorImage.String,
					Following: r.Following,
				},
			},
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount.Int64),
		})
	}
	return results, nil
}

func (r *Articles) GetAllArticlesByAuthor(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	author string,
) ([]domain.ArticleDetail, error) {
	results := []domain.ArticleDetail{}
	rows, err := r.FetchAllArticlesByAuthor(ctx, sqlc.FetchAllArticlesByAuthorParams{
		Username: author,
		Limit:    int32(limit),
		Offset:   int32(offset),
		UserID:   int64(viewerID),
	})
	if err != nil {
		return results, err
	}

	for _, r := range rows {
		results = append(results, domain.ArticleDetail{
			Article: domain.Article{
				ID:          int(r.ID),
				Slug:        r.Slug,
				Title:       r.Title,
				Description: r.Description,
				Body:        r.Body,
				Tags:        r.Tags,
				CreatedAt:   r.CreatedAt.Time,
				UpdatedAt:   r.UpdatedAt.Time,
				Author: domain.Author{
					ID:        int(r.AuthorID),
					Name:      r.AuthorName.String,
					Bio:       r.AuthorBio.String,
					Image:     r.AuthorImage.String,
					Following: r.Following,
				},
			},
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount.Int64),
		})
	}
	return results, nil
}

func (r *Articles) GetAllArticlesByFavorited(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	favoritedUser string,
) ([]domain.ArticleDetail, error) {
	results := []domain.ArticleDetail{}
	rows, err := r.FetchAllArticlesByFavorited(ctx, sqlc.FetchAllArticlesByFavoritedParams{
		Username: favoritedUser,
		Limit:    int32(limit),
		Offset:   int32(offset),
		UserID:   int64(viewerID),
	})
	if err != nil {
		return results, err
	}

	for _, r := range rows {
		results = append(results, domain.ArticleDetail{
			Article: domain.Article{
				ID:          int(r.ID),
				Slug:        r.Slug,
				Title:       r.Title,
				Description: r.Description,
				Body:        r.Body,
				Tags:        r.Tags,
				CreatedAt:   r.CreatedAt.Time,
				UpdatedAt:   r.UpdatedAt.Time,
				Author: domain.Author{
					ID:        int(r.AuthorID),
					Name:      r.AuthorName,
					Bio:       r.AuthorBio.String,
					Image:     r.AuthorImage.String,
					Following: r.Following,
				},
			},
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount.Int64),
		})
	}
	return results, nil
}

func (r *Articles) GetAllArticlesByTag(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	tag string,
) ([]domain.ArticleDetail, error) {
	results := []domain.ArticleDetail{}
	rows, err := r.FetchAllArticlesByTag(ctx, sqlc.FetchAllArticlesByTagParams{
		Name:   tag,
		Limit:  int32(limit),
		Offset: int32(offset),
		UserID: int64(viewerID),
	})
	if err != nil {
		return results, err
	}

	for _, r := range rows {
		results = append(results, domain.ArticleDetail{
			Article: domain.Article{
				ID:          int(r.ID),
				Slug:        r.Slug,
				Title:       r.Title,
				Description: r.Description,
				Body:        r.Body,
				Tags:        r.Tags,
				CreatedAt:   r.CreatedAt.Time,
				UpdatedAt:   r.UpdatedAt.Time,
				Author: domain.Author{
					ID:        int(r.AuthorID),
					Name:      r.AuthorName,
					Bio:       r.AuthorBio.String,
					Image:     r.AuthorImage.String,
					Following: r.Following,
				},
			},
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount.Int64),
		})
	}
	return results, nil
}
