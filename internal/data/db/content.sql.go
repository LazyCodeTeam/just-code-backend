// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: content.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAssetById = `-- name: DeleteAssetById :exec
DELETE FROM asset WHERE id = $1
`

func (q *Queries) DeleteAssetById(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteAssetById, id)
	return err
}

const deleteSectionById = `-- name: DeleteSectionById :exec
DELETE FROM section WHERE id = $1
`

func (q *Queries) DeleteSectionById(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteSectionById, id)
	return err
}

const deleteTaskById = `-- name: DeleteTaskById :exec
DELETE FROM task WHERE id = $1
`

func (q *Queries) DeleteTaskById(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteTaskById, id)
	return err
}

const deleteTechnologyById = `-- name: DeleteTechnologyById :exec
DELETE FROM technology WHERE id = $1
`

func (q *Queries) DeleteTechnologyById(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteTechnologyById, id)
	return err
}

const getAllAssets = `-- name: GetAllAssets :many
SELECT id, url, created_at FROM asset
`

func (q *Queries) GetAllAssets(ctx context.Context) ([]Asset, error) {
	rows, err := q.db.Query(ctx, getAllAssets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asset
	for rows.Next() {
		var i Asset
		if err := rows.Scan(&i.ID, &i.Url, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSectionTasks = `-- name: GetAllSectionTasks :many
SELECT id, section_id, title, description, image_url, difficulty, content, position, is_public, updated_at, created_at FROM task WHERE section_id = $1 AND position IS NOT NULL ORDER BY position ASC
`

func (q *Queries) GetAllSectionTasks(ctx context.Context, sectionID pgtype.UUID) ([]Task, error) {
	rows, err := q.db.Query(ctx, getAllSectionTasks, sectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.SectionID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Difficulty,
			&i.Content,
			&i.Position,
			&i.IsPublic,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSections = `-- name: GetAllSections :many
SELECT id, technology_id, title, description, image_url, position, updated_at, created_at FROM section
`

func (q *Queries) GetAllSections(ctx context.Context) ([]Section, error) {
	rows, err := q.db.Query(ctx, getAllSections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Section
	for rows.Next() {
		var i Section
		if err := rows.Scan(
			&i.ID,
			&i.TechnologyID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Position,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, section_id, title, description, image_url, difficulty, content, position, is_public, updated_at, created_at FROM task
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.Query(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.SectionID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Difficulty,
			&i.Content,
			&i.Position,
			&i.IsPublic,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTechnologies = `-- name: GetAllTechnologies :many
SELECT id, title, description, image_url, position, updated_at, created_at FROM technology ORDER BY position ASC
`

func (q *Queries) GetAllTechnologies(ctx context.Context) ([]Technology, error) {
	rows, err := q.db.Query(ctx, getAllTechnologies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Technology
	for rows.Next() {
		var i Technology
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Position,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTechnologiesWithSectionsPreview = `-- name: GetAllTechnologiesWithSectionsPreview :many
SELECT technology.id, technology.title, technology.description, technology.image_url, technology.position, technology.updated_at, technology.created_at, 
  section.id as section_id, 
  section.title as section_title
FROM technology
JOIN section ON section.technology_id = technology.id
ORDER BY technology.position ASC, section.position ASC
`

type GetAllTechnologiesWithSectionsPreviewRow struct {
	ID           pgtype.UUID
	Title        string
	Description  pgtype.Text
	ImageUrl     pgtype.Text
	Position     int32
	UpdatedAt    pgtype.Timestamptz
	CreatedAt    pgtype.Timestamptz
	SectionID    pgtype.UUID
	SectionTitle string
}

func (q *Queries) GetAllTechnologiesWithSectionsPreview(ctx context.Context) ([]GetAllTechnologiesWithSectionsPreviewRow, error) {
	rows, err := q.db.Query(ctx, getAllTechnologiesWithSectionsPreview)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTechnologiesWithSectionsPreviewRow
	for rows.Next() {
		var i GetAllTechnologiesWithSectionsPreviewRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Position,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.SectionID,
			&i.SectionTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTechnologySections = `-- name: GetAllTechnologySections :many
SELECT id, technology_id, title, description, image_url, position, updated_at, created_at FROM section WHERE technology_id = $1 ORDER BY position ASC
`

func (q *Queries) GetAllTechnologySections(ctx context.Context, technologyID pgtype.UUID) ([]Section, error) {
	rows, err := q.db.Query(ctx, getAllTechnologySections, technologyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Section
	for rows.Next() {
		var i Section
		if err := rows.Scan(
			&i.ID,
			&i.TechnologyID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Position,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTechnolotySectionsWithTasksPreview = `-- name: GetAllTechnolotySectionsWithTasksPreview :many
SELECT section.id, section.technology_id, section.title, section.description, section.image_url, section.position, section.updated_at, section.created_at, 
  task.id as task_id, 
  task.title as task_title, 
  task.is_public as task_is_public
FROM section
JOIN task ON task.section_id = section.id
WHERE section.technology_id = $1 AND task.position IS NOT NULL
ORDER BY section.position ASC, task.position ASC
`

type GetAllTechnolotySectionsWithTasksPreviewRow struct {
	ID           pgtype.UUID
	TechnologyID pgtype.UUID
	Title        string
	Description  pgtype.Text
	ImageUrl     pgtype.Text
	Position     int32
	UpdatedAt    pgtype.Timestamptz
	CreatedAt    pgtype.Timestamptz
	TaskID       pgtype.UUID
	TaskTitle    string
	TaskIsPublic bool
}

func (q *Queries) GetAllTechnolotySectionsWithTasksPreview(ctx context.Context, technologyID pgtype.UUID) ([]GetAllTechnolotySectionsWithTasksPreviewRow, error) {
	rows, err := q.db.Query(ctx, getAllTechnolotySectionsWithTasksPreview, technologyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTechnolotySectionsWithTasksPreviewRow
	for rows.Next() {
		var i GetAllTechnolotySectionsWithTasksPreviewRow
		if err := rows.Scan(
			&i.ID,
			&i.TechnologyID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Position,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.TaskID,
			&i.TaskTitle,
			&i.TaskIsPublic,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAsset = `-- name: InsertAsset :one
INSERT INTO asset (
  id,
  url
) VALUES (
  $1, $2
) RETURNING id, url, created_at
`

type InsertAssetParams struct {
	ID  pgtype.UUID
	Url string
}

func (q *Queries) InsertAsset(ctx context.Context, arg InsertAssetParams) (Asset, error) {
	row := q.db.QueryRow(ctx, insertAsset, arg.ID, arg.Url)
	var i Asset
	err := row.Scan(&i.ID, &i.Url, &i.CreatedAt)
	return i, err
}

const upsertSection = `-- name: UpsertSection :exec
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
  updated_at = NOW()
`

type UpsertSectionParams struct {
	ID           pgtype.UUID
	TechnologyID pgtype.UUID
	Title        string
	Description  pgtype.Text
	ImageUrl     pgtype.Text
	Position     int32
}

func (q *Queries) UpsertSection(ctx context.Context, arg UpsertSectionParams) error {
	_, err := q.db.Exec(ctx, upsertSection,
		arg.ID,
		arg.TechnologyID,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Position,
	)
	return err
}

const upsertTask = `-- name: UpsertTask :exec
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
  updated_at = NOW()
`

type UpsertTaskParams struct {
	ID          pgtype.UUID
	SectionID   pgtype.UUID
	Title       string
	Description pgtype.Text
	ImageUrl    pgtype.Text
	Difficulty  int32
	Content     []byte
	Position    pgtype.Int4
	IsPublic    bool
}

func (q *Queries) UpsertTask(ctx context.Context, arg UpsertTaskParams) error {
	_, err := q.db.Exec(ctx, upsertTask,
		arg.ID,
		arg.SectionID,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Difficulty,
		arg.Content,
		arg.Position,
		arg.IsPublic,
	)
	return err
}

const upsertTechnology = `-- name: UpsertTechnology :exec
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
  updated_at = NOW()
`

type UpsertTechnologyParams struct {
	ID          pgtype.UUID
	Title       string
	Description pgtype.Text
	ImageUrl    pgtype.Text
	Position    int32
}

func (q *Queries) UpsertTechnology(ctx context.Context, arg UpsertTechnologyParams) error {
	_, err := q.db.Exec(ctx, upsertTechnology,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Position,
	)
	return err
}
