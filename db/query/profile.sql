-- name: GetProfileById :one
SELECT * FROM profile
WHERE id = $1 
LIMIT 1;

-- name: CreateProfile :one
INSERT INTO profile (
  id, name, first_name, last_name
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (id) DO UPDATE
SET
  name = $2,
  first_name = $3,
  last_name = $4,
  updated_at = NOW()
RETURNING *;

-- name: UpdateProfileAvatar :exec
UPDATE profile SET
  avatar_url = $2
WHERE id = $1;
