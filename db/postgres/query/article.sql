-- name: InsertArticle :one
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
RETURNING *;

-- name: GetArticleDetail :one
WITH article_cte AS (
    SELECT *
    FROM articles
    WHERE slug=sqlc.arg('slug')
),
author_cte AS (
    SELECT *
    FROM users
    WHERE id=(SELECT author_id FROM article_cte LIMIT 1)
),
favorites_cte AS (
    SELECT COUNT(*) as count
    FROM favorites
    WHERE article_id=(SELECT id FROM article_cte LIMIT 1)
)
SELECT
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    a.created_at,
    a.updated_at,
    f.count AS favorites_count,
    CASE WHEN EXISTS (
        SELECT 1 FROM favorites
        WHERE favorites.article_id = a.id
        AND favorites.user_id = sqlc.arg('viewer_id')
    ) THEN true ELSE false
    END AS favorited,
    CASE WHEN EXISTS (
        SELECT 1 FROM follows
        WHERE follows.following_id = a.author_id
        AND follows.follower_id = sqlc.arg('viewer_id')
    ) THEN true ELSE false
    END AS following,
    author.id AS author_id,
    author.username,
    author.bio,
    author.image
FROM article_cte AS a, author_cte as author, favorites_cte as f;

-- name: UpdateArticle :one
UPDATE articles
SET slug = coalesce(sqlc.narg('slug'), slug),
    title = coalesce(sqlc.narg('title'), title),
    description = coalesce(sqlc.narg('description'), description),
    body = coalesce(sqlc.narg('body'), body),
    updated_at = now()
WHERE id = sqlc.arg('id') AND author_id = sqlc.arg('author_id')
RETURNING *;

-- name: GetArticleBySlug :one
SELECT *
FROM articles
WHERE slug=$1;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;

-- name: InsertFavorite :exec
INSERT INTO favorites (
    user_id,
    article_id
) VALUES (
    $1,
    $2
);

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id=$1 AND article_id=$2;

-- name: InsertTag :one
INSERT INTO tags (name) VALUES ($1)
ON CONFLICT DO NOTHING
RETURNING id;

-- name: InsertArticleTag :exec
INSERT INTO article_tags (article_id, tag_id) VALUES ($1, $2);