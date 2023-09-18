-- name: GetAllTechnologiesWithSectionsPreview :many
SELECT technology.*, 
  section.id as section_id, 
  section.title as section_title
FROM technology
JOIN section ON section.technology_id = technology.id
ORDER BY technology.position ASC, section.position ASC;

-- name: GetAllTechnologies :many
SELECT * FROM technology ORDER BY position ASC;

-- name: GetAllTechnolotySectionsWithTasksPreview :many
SELECT section.*, 
  task.id as task_id, 
  task.title as task_title, 
  task.is_public as task_is_public
FROM section
JOIN task ON task.section_id = section.id
WHERE section.technology_id = $1 AND task.position IS NOT NULL
ORDER BY section.position ASC, task.position ASC;

-- name: GetAllTechnologySections :many
SELECT * FROM section WHERE technology_id = $1 ORDER BY position ASC;

-- name: GetAllSections :many
SELECT * FROM section;

-- name: GetAllSectionTasks :many
SELECT * FROM task WHERE section_id = $1 AND position IS NOT NULL ORDER BY position ASC;

-- name: GetAllTasks :many
SELECT * FROM task;

-- name: UpsertTechnology :exec
INSERT INTO technology (
  id,
  title,
  description,
  image_url,
  position
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (id) DO UPDATE
SET
  title = $2,
  description = $3,
  image_url = $4,
  position = $5,
  updated_at = NOW();

-- name: DeleteTechnologyById :exec
DELETE FROM technology WHERE id = $1;

-- name: UpsertSection :exec
INSERT INTO section (
  id,
  technology_id,
  title,
  description,
  image_url,
  position 
) VALUES (
  $1, $2, $3, $4, $5, $6
)
ON CONFLICT (id) DO UPDATE
SET
  technology_id = $2,
  title = $3,
  description = $4,
  image_url = $5,
  position = $6,
  updated_at = NOW();

-- name: DeleteSectionById :exec
DELETE FROM section WHERE id = $1;

-- name: UpsertTask :exec
INSERT INTO task (
  id,
  section_id,
  title,
  description,
  image_url,
  difficulty,
  content,
  position,
  is_public
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (id) DO UPDATE
SET
  section_id = $2,
  title = $3,
  description = $4,
  image_url = $5,
  difficulty = $6,
  content = $7,
  position = $8,
  is_public = $9,
  updated_at = NOW();

-- name: DeleteTaskById :exec
DELETE FROM task WHERE id = $1;

-- name: DeleteAssetById :exec
DELETE FROM asset WHERE id = $1;

-- name: InsertAsset :one
INSERT INTO asset (
  id,
  url
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetAllAssets :many
SELECT * FROM asset;

-- name: GetTaskById :one
SELECT DISTINCT ON (answer.task_id) 
  task.*, 
  answer.created_at as answer_done_at
FROM task 
LEFT JOIN answer ON answer.task_id = task.id AND answer.result = 'FIRST_VALID'
WHERE task.id = $1
LIMIT 1;
