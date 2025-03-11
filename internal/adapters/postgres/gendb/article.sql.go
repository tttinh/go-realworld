// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: article.sql

package gendb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteArticle = `-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1
`

func (q *Queries) DeleteArticle(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteArticle, id)
	return err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE id=$1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const deleteFavorite = `-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id=$1 AND article_id=$2
`

type DeleteFavoriteParams struct {
	UserID    int64
	ArticleID int64
}

func (q *Queries) DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) error {
	_, err := q.db.Exec(ctx, deleteFavorite, arg.UserID, arg.ArticleID)
	return err
}

const fetchAllComments = `-- name: FetchAllComments :many
SELECT id, body, author_id, article_id, created_at, updated_at
FROM comments
WHERE article_id=$1
`

func (q *Queries) FetchAllComments(ctx context.Context, articleID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, fetchAllComments, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.AuthorID,
			&i.ArticleID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchAllTags = `-- name: FetchAllTags :one
SELECT array_agg(t.name)::text[]
FROM tags t
`

func (q *Queries) FetchAllTags(ctx context.Context) ([]string, error) {
	row := q.db.QueryRow(ctx, fetchAllTags)
	var column_1 []string
	err := row.Scan(&column_1)
	return column_1, err
}

const fetchArticleBySlug = `-- name: FetchArticleBySlug :one
SELECT id, author_id, slug, title, description, body, created_at, updated_at
FROM articles
WHERE slug=$1
`

func (q *Queries) FetchArticleBySlug(ctx context.Context, slug string) (Article, error) {
	row := q.db.QueryRow(ctx, fetchArticleBySlug, slug)
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

const fetchArticleDetail = `-- name: FetchArticleDetail :one
WITH article_cte AS (
    SELECT id, author_id, slug, title, description, body, created_at, updated_at
    FROM articles
    WHERE slug=$2
),
author_cte AS (
    SELECT id, username, email, password, bio, image, created_at, updated_at
    FROM users
    WHERE id=(SELECT author_id FROM article_cte LIMIT 1)
),
favorite_cte AS (
    SELECT COUNT(*) as count
    FROM favorites
    WHERE article_id=(SELECT id FROM article_cte LIMIT 1)
),
tag_cte AS (
	SELECT array_agg(t.name) FILTER (WHERE t.name IS NOT NULL)::text[] AS names
	FROM tags t
	LEFT JOIN article_tags at ON t.id = at.tag_id
	WHERE at.article_id=(SELECT id FROM article_cte LIMIT 1)
)
SELECT
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    a.created_at,
    a.updated_at,
    t.names AS tags,
    f.count AS favorites_count,
    CASE WHEN EXISTS (
        SELECT 1 FROM favorites
        WHERE favorites.article_id = a.id
        AND favorites.user_id = $1
    ) THEN true ELSE false
    END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1 FROM follows
        WHERE follows.following_id = a.author_id
        AND follows.follower_id = $1
    ) THEN true ELSE false
    END AS following,
    author.id AS author_id,
    author.username,
    author.bio,
    author.image
FROM article_cte AS a, author_cte as author, favorite_cte as f, tag_cte as t
`

type FetchArticleDetailParams struct {
	ViewerID int64
	Slug     string
}

type FetchArticleDetailRow struct {
	ID             int64
	Slug           string
	Title          string
	Description    string
	Body           string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	Tags           []string
	FavoritesCount int64
	Favorited      bool
	Following      bool
	AuthorID       int64
	Username       string
	Bio            pgtype.Text
	Image          pgtype.Text
}

func (q *Queries) FetchArticleDetail(ctx context.Context, arg FetchArticleDetailParams) (FetchArticleDetailRow, error) {
	row := q.db.QueryRow(ctx, fetchArticleDetail, arg.ViewerID, arg.Slug)
	var i FetchArticleDetailRow
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.Title,
		&i.Description,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Tags,
		&i.FavoritesCount,
		&i.Favorited,
		&i.Following,
		&i.AuthorID,
		&i.Username,
		&i.Bio,
		&i.Image,
	)
	return i, err
}

const fetchCommentByID = `-- name: FetchCommentByID :one
SELECT id, body, author_id, article_id, created_at, updated_at
FROM comments
WHERE id=$1
`

func (q *Queries) FetchCommentByID(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRow(ctx, fetchCommentByID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.AuthorID,
		&i.ArticleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertArticle = `-- name: InsertArticle :one
INSERT INTO articles (
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
    $5
)
RETURNING id, author_id, slug, title, description, body, created_at, updated_at
`

type InsertArticleParams struct {
	AuthorID    int64
	Slug        string
	Title       string
	Description string
	Body        string
}

func (q *Queries) InsertArticle(ctx context.Context, arg InsertArticleParams) (Article, error) {
	row := q.db.QueryRow(ctx, insertArticle,
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

const insertArticleTag = `-- name: InsertArticleTag :exec
INSERT INTO article_tags (
    article_id,
    tag_id
) VALUES (
    $1,
    $2
)
`

type InsertArticleTagParams struct {
	ArticleID int64
	TagID     int64
}

func (q *Queries) InsertArticleTag(ctx context.Context, arg InsertArticleTagParams) error {
	_, err := q.db.Exec(ctx, insertArticleTag, arg.ArticleID, arg.TagID)
	return err
}

const insertComment = `-- name: InsertComment :one
INSERT INTO comments (
    article_id,
    author_id,
    body
) VALUES (
    $1,
    $2,
    $3
)
RETURNING id, body, author_id, article_id, created_at, updated_at
`

type InsertCommentParams struct {
	ArticleID int64
	AuthorID  int64
	Body      string
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, insertComment, arg.ArticleID, arg.AuthorID, arg.Body)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.AuthorID,
		&i.ArticleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertFavorite = `-- name: InsertFavorite :exec
INSERT INTO favorites (
    user_id,
    article_id
) VALUES (
    $1,
    $2
)
`

type InsertFavoriteParams struct {
	UserID    int64
	ArticleID int64
}

func (q *Queries) InsertFavorite(ctx context.Context, arg InsertFavoriteParams) error {
	_, err := q.db.Exec(ctx, insertFavorite, arg.UserID, arg.ArticleID)
	return err
}

const insertTag = `-- name: InsertTag :one
INSERT INTO tags (
    name
) VALUES (
    $1
)
ON CONFLICT
    ON CONSTRAINT tags_name_key
DO UPDATE SET name = $1
RETURNING id
`

func (q *Queries) InsertTag(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, insertTag, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateArticle = `-- name: UpdateArticle :one
UPDATE articles
SET slug = coalesce($1, slug),
    title = coalesce($2, title),
    description = coalesce($3, description),
    body = coalesce($4, body),
    updated_at = now()
WHERE id = $5 AND author_id = $6
RETURNING id, author_id, slug, title, description, body, created_at, updated_at
`

type UpdateArticleParams struct {
	Slug        pgtype.Text
	Title       pgtype.Text
	Description pgtype.Text
	Body        pgtype.Text
	ID          int64
	AuthorID    int64
}

func (q *Queries) UpdateArticle(ctx context.Context, arg UpdateArticleParams) (Article, error) {
	row := q.db.QueryRow(ctx, updateArticle,
		arg.Slug,
		arg.Title,
		arg.Description,
		arg.Body,
		arg.ID,
		arg.AuthorID,
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

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET body=$2
WHERE id=$1
RETURNING id, body, author_id, article_id, created_at, updated_at
`

type UpdateCommentParams struct {
	ID   int64
	Body string
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, updateComment, arg.ID, arg.Body)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.AuthorID,
		&i.ArticleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
