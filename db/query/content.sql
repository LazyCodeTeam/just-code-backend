-- name: GetAllTechnologiesWithSectionsPreview :many
SELECT technology.*, 
  section.id as section_id, 
  section.position, 
  section.title as section_title, 
  section.description as section_description, 
  section.image_url as section_image_url
FROM technology
LEFT JOIN section ON section.technology_id = technology.id
ORDER BY technology.position ASC, section.position ASC;

-- name: GetAllTechnologies :many
SELECT * FROM technology ORDER BY position ASC;

-- name: GetAllTechnolotySectionsWithTasksPreview :many
SELECT section.*, 
  task.id as task_id, 
  task.position, 
  task.title as task_title, 
  task.image_url as task_image_url
FROM section
LEFT JOIN task ON task.section_id = section.id
WHERE section.technology_id = $1
ORDER BY section.position ASC, task.position ASC;

-- name: GetAllTechnologySections :many
SELECT * FROM section WHERE technology_id = $1 ORDER BY position ASC;

-- name: GetAllSectionTasks :many
SELECT * FROM task WHERE section_id = $1 ORDER BY position ASC;

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

-- name: UpsertTask :exec
INSERT INTO task (
  id,
  section_id,
  title,
  image_url,
  difficulty,
  content,
  position,
  is_dynamic,
  is_public
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (id) DO UPDATE
SET
  section_id = $2,
  title = $3,
  image_url = $4,
  difficulty = $5,
  content = $6,
  position = $7,
  is_dynamic = $8,
  is_public = $9,
  updated_at = NOW();
