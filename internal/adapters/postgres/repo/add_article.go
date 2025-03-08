package pgrepo

import (
	"context"
	"fmt"

	"github.com/tinhtt/go-realworld/internal/adapters/postgres/gendb"
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

func createArticleTx(q *gendb.Queries, ctx context.Context, a domain.Article) (domain.Article, error) {
	param := gendb.InsertArticleParams{
		AuthorID:    int64(a.AuthorID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
	dbArticle, err := q.InsertArticle(ctx, param)
	if err != nil {
		if derr := toDomainError(err); derr != nil {
			return domain.Article{}, derr
		}
		return domain.Article{}, err
	}

	for _, tag := range a.Tags {
		tagID, err := q.InsertTag(ctx, tag)
		if err != nil {
			return domain.Article{}, fmt.Errorf("insert tag: %w", err)
		}

		err = q.InsertArticleTag(ctx, gendb.InsertArticleTagParams{
			ArticleID: dbArticle.ID,
			TagID:     tagID,
		})
		if err != nil {
			return domain.Article{}, fmt.Errorf("insert article tag: %w", err)
		}
	}

	return domain.Article{
		ID:          int(dbArticle.ID),
		AuthorID:    int(dbArticle.AuthorID),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		Tags:        a.Tags,
		CreatedAt:   dbArticle.CreatedAt.Time,
		UpdatedAt:   dbArticle.UpdatedAt.Time,
	}, nil
}
