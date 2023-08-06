// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: profile.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createProfile = `-- name: CreateProfile :one
INSERT INTO profile (
  id, name, avatar_url, first_name, last_name, updated_at, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, name, avatar_url, first_name, last_name, updated_at, created_at
`

type CreateProfileParams struct {
	ID        string
	Name      string
	AvatarUrl sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, createProfile,
		arg.ID,
		arg.Name,
		arg.AvatarUrl,
		arg.FirstName,
		arg.LastName,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.FirstName,
		&i.LastName,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getProfileById = `-- name: GetProfileById :one
SELECT id, name, avatar_url, first_name, last_name, updated_at, created_at FROM profile
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProfileById(ctx context.Context, id string) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileById, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AvatarUrl,
		&i.FirstName,
		&i.LastName,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
