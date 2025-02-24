-- name: CreateArticle :one
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
RETURNING *;

-- name: GetArticle :one
SELECT *
FROM articles
WHERE id=$1;

-- name: GetArticleBySlug :one
SELECT *
FROM articles
WHERE slug=$1;

-- name: UpdateArticle :one
UPDATE articles
SET slug = coalesce(sqlc.narg('slug'), slug),
    title = coalesce(sqlc.narg('title'), title),
    description = coalesce(sqlc.narg('description'), description),
    body = coalesce(sqlc.narg('body'), body),
    updated_at = now()
WHERE id = sqlc.arg('id') AND author_id = sqlc.arg('author_id')
RETURNING *;