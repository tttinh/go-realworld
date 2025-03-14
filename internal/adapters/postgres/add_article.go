package postgres

import (
	"context"
	"fmt"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) Add(ctx context.Context, a domain.Article) (domain.Article, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return domain.Article{}, err
	}
	defer tx.Rollback(ctx)

	qtx := r.WithTx(tx)
	res, err := createArticleTx(qtx, ctx, a)
	if err != nil {
		return domain.Article{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Article{}, err
	}

	return res, nil
}

func createArticleTx(q *sqlc.Queries, ctx context.Context, a domain.Article) (domain.Article, error) {
	param := sqlc.InsertArticleParams{
		AuthorID:    int64(a.Author.ID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
	dbArticle, err := q.InsertArticle(ctx, param)
	if err != nil {
		return domain.Article{}, toDomainError(err)
	}

	for _, tag := range a.Tags {
		tagID, err := q.InsertTag(ctx, tag)
		if err != nil {
			return domain.Article{}, fmt.Errorf("insert tag: %w", err)
		}

		err = q.InsertArticleTag(ctx, sqlc.InsertArticleTagParams{
			ArticleID: dbArticle.ID,
			TagID:     tagID,
		})
		if err != nil {
			return domain.Article{}, fmt.Errorf("insert article tag: %w", err)
		}
	}

	return domain.Article{
		ID:          int(dbArticle.ID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		Tags:        a.Tags,
		CreatedAt:   dbArticle.CreatedAt.Time,
		UpdatedAt:   dbArticle.UpdatedAt.Time,
		Author:      domain.Author{ID: int(dbArticle.AuthorID)},
	}, nil
}
