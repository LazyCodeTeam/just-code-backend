package port

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

type ContentRepository interface {
	UpsertTask(ctx context.Context, task model.Task) error

	UpsertSection(ctx context.Context, section model.Section) error

	UpsertTechnology(ctx context.Context, technology model.Technology) error

	GetAllTasks(ctx context.Context) ([]model.Task, error)

	GetAllSections(ctx context.Context) ([]model.Section, error)

	GetAllTechnologies(ctx context.Context) ([]model.Technology, error)

	DeleteTaskById(ctx context.Context, id string) error

	DeleteSectionById(ctx context.Context, id string) error

	DeleteTechnologyById(ctx context.Context, id string) error

	GetTechnologiesWithSectionsPreview(
		ctx context.Context,
	) ([]model.TechnologyWithSectionsPreview, error)

	GetSectionsWithTasksPreview(
		ctx context.Context,
		technologyID string,
	) ([]model.SectionWithTasksPreview, error)

	SaveAsset(ctx context.Context, id string, url string) (model.Asset, error)

	DeleteAsset(ctx context.Context, id string) error

	GetAssets(ctx context.Context) ([]model.Asset, error)

	GetSectionTasks(ctx context.Context, sectionID string) ([]model.Task, error)
}
