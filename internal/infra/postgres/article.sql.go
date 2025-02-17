// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: article.sql

package pgdb

import (
	"context"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (
    id,
    author_id,
    slug,
    title,
    description,
    body
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, author_id, slug, title, description, body, created_at, updated_at
`

type CreateArticleParams struct {
	ID          int64
	AuthorID    int64
	Slug        string
	Title       string
	Description string
	Body        string
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRow(ctx, createArticle,
		arg.ID,
		arg.AuthorID,
		arg.Slug,
		arg.Title,
		arg.Description,
		arg.Body,
	)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
