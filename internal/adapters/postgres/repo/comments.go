package pgrepo

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/gendb"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetAllComments(ctx context.Context, articleID int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	rows, err := r.FetchAllComments(ctx, int64(articleID))
	if err != nil {
		return comments, err
	}

	for _, r := range rows {
		comments = append(comments, domain.Comment{
			ID:        int(r.ID),
			AuthorID:  int(r.AuthorID),
			ArticleID: int(r.ArticleID),
			CreatedAt: r.CreatedAt.Time,
			UpdatedAt: r.UpdatedAt.Time,
		})
	}
	return comments, nil
}

func (r *Articles) GetComment(ctx context.Context, id int) (domain.Comment, error) {
	row, err := r.FetchCommentByID(ctx, int64(id))
	if err != nil {
		return domain.Comment{}, err
	}

	return domain.Comment{
		ID:        int(row.ID),
		AuthorID:  int(row.AuthorID),
		ArticleID: int(row.ArticleID),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *Articles) AddComment(ctx context.Context, c domain.Comment) (domain.Comment, error) {
	row, err := r.InsertComment(ctx, gendb.InsertCommentParams{
		ArticleID: int64(c.ArticleID),
		AuthorID:  int64(c.AuthorID),
		Body:      c.Body,
	})
	if err != nil {
		return domain.Comment{}, err
	}

	return domain.Comment{
		ID:        int(row.ID),
		AuthorID:  int(row.AuthorID),
		ArticleID: int(row.ArticleID),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *Articles) RemoveComment(ctx context.Context, id int) error {
	return r.DeleteComment(ctx, int64(id))
}
