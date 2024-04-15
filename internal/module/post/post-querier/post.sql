-- name: CreateOne :one
INSERT INTO posts (title, slug, category, status, content, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING *;

-- name: GetOneById :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetMany :many
SELECT *
FROM posts
WHERE
  (
    (sqlc.narg('title_eq')::text IS NOT NULL AND title = sqlc.narg('title_eq'))
    OR
    (sqlc.narg('title_regex')::text IS NOT NULL AND title ~* sqlc.narg('title_regex')::text)
    OR
    (sqlc.narg('title_eq')::text IS NULL AND sqlc.narg('title_regex') IS NULL)
  )
  AND
  (
    (sqlc.narg('slug_eq')::text IS NOT NULL AND slug = sqlc.narg('slug_eq'))
    OR
    (sqlc.narg('slug_regex')::text IS NOT NULL AND slug ~* sqlc.narg('slug_regex')::text)
    OR
    (sqlc.narg('slug_eq')::text IS NULL AND sqlc.narg('slug_regex') IS NULL)
  )
  AND
  (
    (sqlc.narg('category_eq')::uuid IS NOT NULL AND category = sqlc.narg('category_eq'))    
    OR
    (sqlc.narg('category_in')::uuid[] IS NOT NULL AND category = ANY(sqlc.narg('category_in')::uuid[]))
    OR
    (sqlc.narg('category_eq')::text IS NULL AND sqlc.narg('category_in') IS NULL)
  )
  AND
  (
    status = sqlc.narg('status') OR sqlc.narg('status') IS NULL
  )
ORDER BY
  CASE WHEN sqlc.narg('sort_query')::text = 'title__asc' THEN title END ASC,
  CASE WHEN sqlc.narg('sort_query')::text = 'title__desc' THEN title END DESC,
  CASE WHEN sqlc.narg('sort_query')::text = 'created_at__asc' THEN created_at END ASC,
  CASE WHEN sqlc.narg('sort_query')::text = 'created_at__desc' THEN created_at END DESC,
  CASE WHEN sqlc.narg('sort_query')::text = 'updated_at__asc' THEN updated_at END ASC,
  CASE WHEN sqlc.narg('sort_query')::text = 'updated_at__desc' THEN updated_at END DESC,
  CASE WHEN sqlc.narg('sort_query')::text = 'status__asc' THEN status END ASC,
  CASE WHEN sqlc.narg('sort_query')::text = 'status__desc' THEN status END DESC
LIMIT sqlc.narg('limit') OFFSET sqlc.narg('offset');

-- name: GetManyByIds :many
SELECT * FROM posts
WHERE id = ANY($1::uuid[]);

-- name: UpdateOneById :one
UPDATE posts SET
  title = coalesce(sqlc.narg(title), title),
  slug = coalesce(sqlc.narg(slug), slug),
  category = coalesce(sqlc.narg(category), category),
  status = coalesce(sqlc.narg(status), status),
  content = coalesce(sqlc.narg(content), content),
  updated_at = now()
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteOne :exec
DELETE FROM posts
WHERE id = $1;
