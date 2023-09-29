-- name: CreatePost :one
INSERT INTO posts (
  id,
  post,
  created_at,
  created_by,
  updated_at,
  updated_by
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1;
