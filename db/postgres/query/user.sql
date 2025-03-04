-- name: InsertUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users
SET username = coalesce(sqlc.narg('username'), username),
    email = coalesce(sqlc.narg('email'), email),
    password = coalesce(sqlc.narg('password'), password),
    bio = coalesce(sqlc.narg('bio'), bio),
    image = coalesce(sqlc.narg('image'), image),
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: InsertFollow :exec
INSERT INTO follows (
    follower_id,
    following_id
) VALUES (
    $1,
    $2
) ON CONFLICT DO NOTHING;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_id=$1 AND following_id=$2;

-- name: GetProfileByName :one
SELECT *,
    CASE WHEN EXISTS (
        SELECT 1 FROM follows f
        WHERE f.follower_id = a.sqlc.arg('follower_id')
        AND f.following_id = users.id
    ) THEN true ELSE false END AS following
FROM users
WHERE username = sqlc.arg('following_name');