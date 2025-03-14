// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: article.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countAllArticles = `-- name: CountAllArticles :one
SELECT COUNT(*) FROM articles
`

func (q *Queries) CountAllArticles(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAllArticles)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countArticlesByAuthor = `-- name: CountArticlesByAuthor :one
SELECT
    COUNT(*)
FROM
    articles a
JOIN users u ON a.author_id=u.id
WHERE u.username=$1
`

func (q *Queries) CountArticlesByAuthor(ctx context.Context, username string) (int64, error) {
	row := q.db.QueryRow(ctx, countArticlesByAuthor, username)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countArticlesByFavorited = `-- name: CountArticlesByFavorited :one
SELECT
    COUNT(*)
FROM
    favorites f
JOIN users u ON f.user_id=u.id
WHERE u.username=$1
`

func (q *Queries) CountArticlesByFavorited(ctx context.Context, username string) (int64, error) {
	row := q.db.QueryRow(ctx, countArticlesByFavorited, username)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countArticlesByTag = `-- name: CountArticlesByTag :one
SELECT
    COUNT(*)
FROM
    article_tags atg
JOIN tags t ON atg.tag_id=t.id
WHERE t.name=$1
`

func (q *Queries) CountArticlesByTag(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, countArticlesByTag, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countFeed = `-- name: CountFeed :one
WITH following_cte AS (
    SELECT f.following_id AS id FROM follows f WHERE f.follower_id=$1
)
SELECT
    COUNT(*)
FROM
    articles a
WHERE a.author_id=ANY(
    SELECT f.following_id
    FROM follows f
    WHERE f.follower_id=$1
)
`

func (q *Queries) CountFeed(ctx context.Context, followerID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countFeed, followerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

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

const fetchAllArticles = `-- name: FetchAllArticles :many
WITH article_cte AS (
    SELECT
        a.id,
        a.author_id,
        a.slug,
        a.title,
        a.body,
        a.description,
        a.created_at,
        a.updated_at,
        u.username AS author_name,
        u.bio AS author_bio,
        u.image AS author_image
    FROM articles a
    JOIN users u ON a.author_id=u.id
    ORDER BY a.updated_at DESC
    LIMIT $1 OFFSET $2
),
tag_cte AS (
    SELECT
        atags.article_id,
        array_agg(t.name)::text[] AS names
    FROM tags t
    LEFT JOIN article_tags atags on atags.tag_id=t.id
    WHERE atags.article_id =  ANY(SELECT id FROM article_cte)
    GROUP BY atags.article_id
),
favorite_cte AS (
    SELECT
        f.article_id,
        COUNT(*)
    FROM favorites f
    WHERE f.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY f.article_id
)
SELECT
    a.id, a.author_id, a.slug, a.title, a.body, a.description, a.created_at, a.updated_at, a.author_name, a.author_bio, a.author_image,
    t.names as tags,
    f.count as favorites_count,
    CASE WHEN EXISTS (
        SELECT 1
        FROM favorites
        WHERE a.id = favorites.article_id AND favorites.user_id=$3
    ) THEN true ELSE false END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1
        FROM follows
        WHERE a.author_id = follows.following_id AND follows.follower_id=$3
    ) THEN true ELSE false END AS following
FROM article_cte a
LEFT JOIN tag_cte t ON a.id=t.article_id
LEFT JOIN favorite_cte f ON a.id=f.article_id
ORDER BY a.updated_at DESC
`

type FetchAllArticlesParams struct {
	Limit  int32
	Offset int32
	UserID int64
}

type FetchAllArticlesRow struct {
	ID             int64
	AuthorID       int64
	Slug           string
	Title          string
	Body           string
	Description    string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	AuthorName     string
	AuthorBio      pgtype.Text
	AuthorImage    pgtype.Text
	Tags           []string
	FavoritesCount pgtype.Int8
	Favorited      bool
	Following      bool
}

func (q *Queries) FetchAllArticles(ctx context.Context, arg FetchAllArticlesParams) ([]FetchAllArticlesRow, error) {
	rows, err := q.db.Query(ctx, fetchAllArticles, arg.Limit, arg.Offset, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAllArticlesRow
	for rows.Next() {
		var i FetchAllArticlesRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Tags,
			&i.FavoritesCount,
			&i.Favorited,
			&i.Following,
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

const fetchAllArticlesByAuthor = `-- name: FetchAllArticlesByAuthor :many
WITH author_cte AS (
    SELECT
        u.id,
        u.username,
        u.bio,
        u.image
    FROM
        users u
    WHERE
        u.username=$1
),
article_cte AS (
    SELECT
        a.id,
        a.author_id,
        a.slug,
        a.title,
        a.body,
        a.description,
        a.created_at,
        a.updated_at
    FROM articles a
    WHERE a.author_id=(SELECT id FROM author_cte LIMIT 1)
    ORDER BY a.updated_at DESC
    LIMIT $2 OFFSET $3
),
tag_cte AS (
    SELECT
        atags.article_id,
        array_agg(t.name)::text[] AS names
    FROM tags t
    LEFT JOIN article_tags atags on atags.tag_id=t.id
    WHERE atags.article_id =  ANY(SELECT id FROM article_cte)
    GROUP BY atags.article_id
),
favorite_cte AS (
    SELECT
        f.article_id,
        COUNT(*)
    FROM favorites f
    WHERE f.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY f.article_id
)
SELECT
    a.id, a.author_id, a.slug, a.title, a.body, a.description, a.created_at, a.updated_at,
    auth.username AS author_name,
    auth.bio AS author_bio,
    auth.image AS author_image,
    t.names as tags,
    f.count as favorites_count,
    CASE WHEN EXISTS (
        SELECT 1
        FROM favorites
        WHERE a.id = favorites.article_id AND favorites.user_id=$4
    ) THEN true ELSE false END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1
        FROM follows
        WHERE a.author_id = follows.following_id AND follows.follower_id=$4
    ) THEN true ELSE false END AS following
FROM article_cte a
JOIN author_cte auth ON auth.id=a.author_id
LEFT JOIN tag_cte t ON a.id=t.article_id
LEFT JOIN favorite_cte f ON a.id=f.article_id
ORDER BY a.updated_at DESC
`

type FetchAllArticlesByAuthorParams struct {
	Username string
	Limit    int32
	Offset   int32
	UserID   int64
}

type FetchAllArticlesByAuthorRow struct {
	ID             int64
	AuthorID       int64
	Slug           string
	Title          string
	Body           string
	Description    string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	AuthorName     string
	AuthorBio      pgtype.Text
	AuthorImage    pgtype.Text
	Tags           []string
	FavoritesCount pgtype.Int8
	Favorited      bool
	Following      bool
}

func (q *Queries) FetchAllArticlesByAuthor(ctx context.Context, arg FetchAllArticlesByAuthorParams) ([]FetchAllArticlesByAuthorRow, error) {
	rows, err := q.db.Query(ctx, fetchAllArticlesByAuthor,
		arg.Username,
		arg.Limit,
		arg.Offset,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAllArticlesByAuthorRow
	for rows.Next() {
		var i FetchAllArticlesByAuthorRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Tags,
			&i.FavoritesCount,
			&i.Favorited,
			&i.Following,
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

const fetchAllArticlesByFavorited = `-- name: FetchAllArticlesByFavorited :many
WITH user_cte AS (
    SELECT
        u.id,
        f.article_id
    FROM
        users u
    LEFT JOIN favorites f ON u.id=f.user_id
    WHERE
        u.username=$1
),
article_cte AS (
    SELECT
        a.id,
        a.author_id,
        a.slug,
        a.title,
        a.body,
        a.description,
        a.created_at,
        a.updated_at,
        u.username AS author_name,
        u.bio AS author_bio,
        u.image AS author_image
    FROM articles a
    JOIN users u ON a.author_id=u.id
    WHERE a.id=ANY(SELECT article_id FROM user_cte)
    ORDER BY a.updated_at DESC
    LIMIT $2 OFFSET $3
),
tag_cte AS (
    SELECT
        atags.article_id,
        array_agg(t.name)::text[] AS names
    FROM tags t
    LEFT JOIN article_tags atags on atags.tag_id=t.id
    WHERE atags.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY atags.article_id
),
favorite_cte AS (
    SELECT
        f.article_id,
        COUNT(*)
    FROM favorites f
    WHERE f.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY f.article_id
)
SELECT
    a.id, a.author_id, a.slug, a.title, a.body, a.description, a.created_at, a.updated_at, a.author_name, a.author_bio, a.author_image,
    t.names as tags,
    f.count as favorites_count,
    CASE WHEN EXISTS (
        SELECT 1
        FROM favorites
        WHERE a.id = favorites.article_id AND favorites.user_id=$4
    ) THEN true ELSE false END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1
        FROM follows
        WHERE a.author_id = follows.following_id AND follows.follower_id=$4
    ) THEN true ELSE false END AS following
FROM article_cte a
LEFT JOIN tag_cte t ON a.id=t.article_id
LEFT JOIN favorite_cte f ON a.id=f.article_id
ORDER BY a.updated_at DESC
`

type FetchAllArticlesByFavoritedParams struct {
	Username string
	Limit    int32
	Offset   int32
	UserID   int64
}

type FetchAllArticlesByFavoritedRow struct {
	ID             int64
	AuthorID       int64
	Slug           string
	Title          string
	Body           string
	Description    string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	AuthorName     string
	AuthorBio      pgtype.Text
	AuthorImage    pgtype.Text
	Tags           []string
	FavoritesCount pgtype.Int8
	Favorited      bool
	Following      bool
}

func (q *Queries) FetchAllArticlesByFavorited(ctx context.Context, arg FetchAllArticlesByFavoritedParams) ([]FetchAllArticlesByFavoritedRow, error) {
	rows, err := q.db.Query(ctx, fetchAllArticlesByFavorited,
		arg.Username,
		arg.Limit,
		arg.Offset,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAllArticlesByFavoritedRow
	for rows.Next() {
		var i FetchAllArticlesByFavoritedRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Tags,
			&i.FavoritesCount,
			&i.Favorited,
			&i.Following,
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

const fetchAllArticlesByTag = `-- name: FetchAllArticlesByTag :many
WITH article_tags_cte AS (
    SELECT
        atag.article_id,
        atag.tag_id
    FROM
        article_tags atag
    JOIN tags t ON t.id=atag.tag_id
    WHERE
        t.name=$1
),
article_cte AS (
    SELECT
        a.id,
        a.author_id,
        a.slug,
        a.title,
        a.body,
        a.description,
        a.created_at,
        a.updated_at,
        u.username AS author_name,
        u.bio AS author_bio,
        u.image AS author_image
    FROM articles a
    JOIN users u ON a.author_id=u.id
    WHERE a.id=ANY(SELECT article_id FROM article_tags_cte)
    ORDER BY a.updated_at DESC
    LIMIT $2 OFFSET $3
),
tag_cte AS (
    SELECT
        atags.article_id,
        array_agg(t.name)::text[] AS names
    FROM tags t
    LEFT JOIN article_tags atags on atags.tag_id=t.id
    WHERE atags.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY atags.article_id
),
favorite_cte AS (
    SELECT
        f.article_id,
        COUNT(*)
    FROM favorites f
    WHERE f.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY f.article_id
)
SELECT
    a.id, a.author_id, a.slug, a.title, a.body, a.description, a.created_at, a.updated_at, a.author_name, a.author_bio, a.author_image,
    t.names as tags,
    f.count as favorites_count,
    CASE WHEN EXISTS (
        SELECT 1
        FROM favorites
        WHERE a.id = favorites.article_id AND favorites.user_id=$4
    ) THEN true ELSE false END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1
        FROM follows
        WHERE a.author_id = follows.following_id AND follows.follower_id=$4
    ) THEN true ELSE false END AS following
FROM article_cte a
LEFT JOIN tag_cte t ON a.id=t.article_id
LEFT JOIN favorite_cte f ON a.id=f.article_id
ORDER BY a.updated_at DESC
`

type FetchAllArticlesByTagParams struct {
	Name   string
	Limit  int32
	Offset int32
	UserID int64
}

type FetchAllArticlesByTagRow struct {
	ID             int64
	AuthorID       int64
	Slug           string
	Title          string
	Body           string
	Description    string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	AuthorName     string
	AuthorBio      pgtype.Text
	AuthorImage    pgtype.Text
	Tags           []string
	FavoritesCount pgtype.Int8
	Favorited      bool
	Following      bool
}

func (q *Queries) FetchAllArticlesByTag(ctx context.Context, arg FetchAllArticlesByTagParams) ([]FetchAllArticlesByTagRow, error) {
	rows, err := q.db.Query(ctx, fetchAllArticlesByTag,
		arg.Name,
		arg.Limit,
		arg.Offset,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAllArticlesByTagRow
	for rows.Next() {
		var i FetchAllArticlesByTagRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Tags,
			&i.FavoritesCount,
			&i.Favorited,
			&i.Following,
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

const fetchAllComments = `-- name: FetchAllComments :many
SELECT
    c.id,
    c.body,
    c.article_id,
    c.created_at,
    c.updated_at,
    u.id::bigint AS author_id,
    u.username AS author_name,
    u.bio AS author_bio,
    u.image AS author_image,
    CASE WHEN EXISTS (
        SELECT 1 FROM follows
        WHERE follows.following_id = c.author_id
        AND follows.follower_id = $2
    ) THEN true ELSE false
    END AS following
FROM comments c
LEFT JOIN users u ON c.author_id=u.id
WHERE c.article_id=$1
`

type FetchAllCommentsParams struct {
	ArticleID int64
	ViewerID  int64
}

type FetchAllCommentsRow struct {
	ID          int64
	Body        string
	ArticleID   int64
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	AuthorID    int64
	AuthorName  pgtype.Text
	AuthorBio   pgtype.Text
	AuthorImage pgtype.Text
	Following   bool
}

func (q *Queries) FetchAllComments(ctx context.Context, arg FetchAllCommentsParams) ([]FetchAllCommentsRow, error) {
	rows, err := q.db.Query(ctx, fetchAllComments, arg.ArticleID, arg.ViewerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAllCommentsRow
	for rows.Next() {
		var i FetchAllCommentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.ArticleID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorID,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Following,
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
SELECT
    c.id,
    c.body,
    c.article_id,
    c.created_at,
    c.updated_at,
    u.id::bigint AS author_id,
    u.username AS author_name,
    u.bio AS author_bio,
    u.image AS author_image,
    CASE WHEN EXISTS (
        SELECT 1 FROM follows
        WHERE follows.following_id = c.author_id
        AND follows.follower_id = $2
    ) THEN true ELSE false
    END AS following
FROM comments c
LEFT JOIN users u ON c.author_id=u.id
WHERE c.id=$1
`

type FetchCommentByIDParams struct {
	ID       int64
	ViewerID int64
}

type FetchCommentByIDRow struct {
	ID          int64
	Body        string
	ArticleID   int64
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	AuthorID    int64
	AuthorName  pgtype.Text
	AuthorBio   pgtype.Text
	AuthorImage pgtype.Text
	Following   bool
}

func (q *Queries) FetchCommentByID(ctx context.Context, arg FetchCommentByIDParams) (FetchCommentByIDRow, error) {
	row := q.db.QueryRow(ctx, fetchCommentByID, arg.ID, arg.ViewerID)
	var i FetchCommentByIDRow
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.ArticleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AuthorID,
		&i.AuthorName,
		&i.AuthorBio,
		&i.AuthorImage,
		&i.Following,
	)
	return i, err
}

const fetchFeed = `-- name: FetchFeed :many
WITH following_cte AS (
    SELECT follower_id, following_id
    FROM follows
    WHERE follower_id=$1
),
article_cte AS (
    SELECT
        a.id,
        a.author_id,
        a.slug,
        a.title,
        a.body,
        a.description,
        a.created_at,
        a.updated_at,
        u.username AS author_name,
        u.bio AS author_bio,
        u.image AS author_image
    FROM articles a
    JOIN users u ON a.author_id=u.id
    WHERE a.author_id=ANY(SELECT f.following_id FROM following_cte f)
    ORDER BY a.updated_at DESC
    OFFSET $2 LIMIT $3
),
tag_cte AS (
    SELECT
        atg.article_id,
        array_agg(t.name)::text[] AS names
    FROM article_tags atg
    JOIN tags t ON atg.tag_id=t.id
    WHERE atg.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY atg.article_id
),
favorite_cte AS (
    SELECT
        f.article_id,
        COUNT(*)
    FROM favorites f
    WHERE f.article_id=ANY(SELECT id FROM article_cte)
    GROUP BY f.article_id
)
SELECT
    a.id, a.author_id, a.slug, a.title, a.body, a.description, a.created_at, a.updated_at, a.author_name, a.author_bio, a.author_image,
    t.names AS tags,
    f.count AS favorites_count,
    CASE WHEN EXISTS (
        SELECT 1 FROM favorites f WHERE f.user_id=$1 AND f.article_id=a.id
    ) THEN true ELSE false END AS favorited
FROM article_cte a
LEFT JOIN tag_cte t ON a.id=t.article_id
LEFT JOIN favorite_cte f ON a.id=f.article_id
ORDER BY a.updated_at DESC
`

type FetchFeedParams struct {
	UserID int64
	Offset int32
	Limit  int32
}

type FetchFeedRow struct {
	ID             int64
	AuthorID       int64
	Slug           string
	Title          string
	Body           string
	Description    string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	AuthorName     string
	AuthorBio      pgtype.Text
	AuthorImage    pgtype.Text
	Tags           []string
	FavoritesCount pgtype.Int8
	Favorited      bool
}

func (q *Queries) FetchFeed(ctx context.Context, arg FetchFeedParams) ([]FetchFeedRow, error) {
	rows, err := q.db.Query(ctx, fetchFeed, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchFeedRow
	for rows.Next() {
		var i FetchFeedRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.AuthorBio,
			&i.AuthorImage,
			&i.Tags,
			&i.FavoritesCount,
			&i.Favorited,
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
