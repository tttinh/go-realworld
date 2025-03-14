package postgres

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetFeed(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
) (domain.ArticleList, error) {
	al := domain.ArticleList{
		Total:    0,
		Articles: []domain.ArticleDetail{},
	}

	count, err := r.CountFeed(ctx, int64(viewerID))
	if err != nil {
		return al, err
	}
	al.Total = int(count)

	rows, err := r.FetchFeed(ctx, sqlc.FetchFeedParams{
		UserID: int64(viewerID),
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return al, err
	}

	for _, r := range rows {
		al.Articles = append(al.Articles, domain.ArticleDetail{
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
					Following: true,
				},
			},
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount.Int64),
		})
	}
	return al, nil
}
