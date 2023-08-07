-- name: GetProfileById :one
SELECT * FROM profile
WHERE id = $1 LIMIT 1;

-- name: CreateProfile :one
INSERT INTO profile (
  id, name, avatar_url, first_name, last_name
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (id) DO UPDATE
SET
  name = $2,
  avatar_url = $3,
  first_name = $4,
  last_name = $5,
  updated_at = NOW()
RETURNING *;
