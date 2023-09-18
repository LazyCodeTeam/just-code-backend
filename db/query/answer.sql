-- name: InsertAnswer :exec
INSERT INTO answer (
  task_id,
  profile_id,
  result
) VALUES (
  $1, $2, $3
);
