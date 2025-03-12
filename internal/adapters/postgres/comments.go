package postgres

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetAllComments(ctx context.Context, viewerID int, articleID int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	rows, err := r.FetchAllComments(ctx, sqlc.FetchAllCommentsParams{
		ArticleID: int64(articleID),
		ViewerID:  int64(viewerID),
	})
	if err != nil {
		return comments, err
	}

	for _, r := range rows {
		comments = append(comments, domain.Comment{
			ID:        int(r.ID),
			ArticleID: int(r.ArticleID),
			CreatedAt: r.CreatedAt.Time,
			UpdatedAt: r.UpdatedAt.Time,
			Author: domain.Author{
				ID:        int(r.AuthorID),
				Name:      r.AuthorName.String,
				Bio:       r.AuthorBio.String,
				Image:     r.AuthorImage.String,
				Following: r.Following,
			},
		})
	}
	return comments, nil
}

func (r *Articles) GetComment(ctx context.Context, viewerID int, commentID int) (domain.Comment, error) {
	row, err := r.FetchCommentByID(ctx, sqlc.FetchCommentByIDParams{
		ID:       int64(commentID),
		ViewerID: int64(viewerID),
	})
	if err != nil {
		return domain.Comment{}, err
	}

	return domain.Comment{
		ID:        int(row.ID),
		ArticleID: int(row.ArticleID),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		Author: domain.Author{
			ID:        int(row.AuthorID),
			Name:      row.AuthorName.String,
			Bio:       row.AuthorBio.String,
			Image:     row.AuthorImage.String,
			Following: row.Following,
		},
	}, nil
}

func (r *Articles) AddComment(ctx context.Context, c domain.Comment) (domain.Comment, error) {
	row, err := r.InsertComment(ctx, sqlc.InsertCommentParams{
		ArticleID: int64(c.ArticleID),
		AuthorID:  int64(c.Author.ID),
		Body:      c.Body,
	})
	if err != nil {
		return domain.Comment{}, err
	}

	return domain.Comment{
		ID:        int(row.ID),
		ArticleID: int(row.ArticleID),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		Author:    domain.Author{ID: int(row.AuthorID)},
	}, nil
}

func (r *Articles) RemoveComment(ctx context.Context, id int) error {
	return r.DeleteComment(ctx, int64(id))
}
