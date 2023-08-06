-- name: GetProfileById :one
SELECT * FROM profile
WHERE id = $1 LIMIT 1;

-- name: CreateProfile :one
INSERT INTO profile (
  id, name, avatar_url, first_name, last_name, updated_at, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;
