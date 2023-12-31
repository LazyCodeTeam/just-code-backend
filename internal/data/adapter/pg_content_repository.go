package adapter

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	coreUtil "github.com/LazyCodeTeam/just-code-backend/internal/core/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/mapper"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

type PgContentRepository struct {
	queries *db.Queries
}

func NewPgContentRepository(queries *db.Queries) *PgContentRepository {
	return &PgContentRepository{queries: queries}
}

func (r *PgContentRepository) UpsertTask(ctx context.Context, task model.Task) error {
	dbTask, err := mapper.UpsertTaskParamsFromDomain(task)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to map task to db", "err", err, "task", task)
		return err
	}

	err = r.queries.UpsertTask(ctx, dbTask)

	if err != nil {
		slog.ErrorContext(ctx, "Failed to upsert task", "err", err, "task", task)
		return err
	}

	return nil
}

func (r *PgContentRepository) UpsertSection(ctx context.Context, section model.Section) error {
	dbSection := mapper.UpsertSectionParamsFromDomain(section)
	err := r.queries.UpsertSection(ctx, dbSection)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to upsert section", "err", err, "section", section)
		return err
	}

	return nil
}

func (r *PgContentRepository) UpsertTechnology(
	ctx context.Context,
	technology model.Technology,
) error {
	dbTechnology := mapper.UpsertTechnologyParamsFromDomain(technology)
	err := r.queries.UpsertTechnology(ctx, dbTechnology)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to upsert technology", "err", err, "technology", technology)
		return err
	}

	return nil
}

func (r *PgContentRepository) GetAllTasks(
	ctx context.Context,
) ([]model.Task, error) {
	dbTasks, err := r.queries.GetAllTasks(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get all tasks", "err", err)
		return nil, err
	}
	tasks, err := coreUtil.TryMapSlice(dbTasks, mapper.TaskToDomain)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to map tasks", "err", err)
		return nil, err
	}

	return tasks, nil
}

func (r *PgContentRepository) GetAllSections(
	ctx context.Context,
) ([]model.Section, error) {
	sections, err := r.queries.GetAllSections(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"Failed to get all sections",
			"err",
			err,
		)
		return nil, err
	}

	return coreUtil.MapSlice(sections, mapper.SectionToDomain), nil
}

func (r *PgContentRepository) GetAllTechnologies(ctx context.Context) ([]model.Technology, error) {
	technologies, err := r.queries.GetAllTechnologies(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get all technologies", "err", err)
		return nil, err
	}

	return coreUtil.MapSlice(technologies, mapper.TechnologyToDomain), nil
}

func (r *PgContentRepository) DeleteTaskById(ctx context.Context, id string) error {
	err := r.queries.DeleteTaskById(ctx, util.ToPgUUID(id))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete task by id", "err", err)
		return err
	}

	return nil
}

func (r *PgContentRepository) DeleteSectionById(ctx context.Context, id string) error {
	err := r.queries.DeleteSectionById(ctx, util.ToPgUUID(id))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete section by id", "err", err)
		return err
	}

	return nil
}

func (r *PgContentRepository) DeleteTechnologyById(ctx context.Context, id string) error {
	err := r.queries.DeleteTechnologyById(ctx, util.ToPgUUID(id))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete technology by id", "err", err)
		return err
	}

	return nil
}

func (r *PgContentRepository) GetTechnologiesWithSectionsPreview(
	ctx context.Context,
) ([]model.TechnologyWithSectionsPreview, error) {
	rows, err := r.queries.GetAllTechnologiesWithSectionsPreview(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get technologies with sections preview", "err", err)
		return nil, err
	}
	technologies := mapper.GetAllTechnologiesWithSectionsPreviewRowsToDomain(rows)

	return technologies, nil
}

func (r *PgContentRepository) GetSectionsWithTasksPreview(
	ctx context.Context,
	technologyID string,
) ([]model.SectionWithTasksPreview, error) {
	rows, err := r.queries.GetAllTechnolotySectionsWithTasksPreview(
		ctx,
		util.ToPgUUID(technologyID),
	)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get sections with tasks preview", "err", err)
		return nil, err
	}
	sections := mapper.GetAllTechnolotySectionsWithTasksPreviewRowsToDomain(rows)

	return sections, nil
}

func (r *PgContentRepository) GetSectionTasks(
	ctx context.Context,
	sectionID string,
) ([]model.Task, error) {
	rows, err := r.queries.GetAllSectionTasks(ctx, util.ToPgUUID(sectionID))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get tasks", "err", err)
		return nil, err
	}

	tasks, err := coreUtil.TryMapSlice(rows, mapper.GetAllSectionTasksRowToDomain)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to map tasks", "err", err)

		return nil, err
	}

	return tasks, nil
}

func (r *PgContentRepository) SaveAsset(
	ctx context.Context,
	id string,
	url string,
) (model.Asset, error) {
	dbAsset, err := r.queries.InsertAsset(ctx, db.InsertAssetParams{
		ID:  util.ToPgUUID(id),
		Url: url,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to save asset", "err", err)

		return model.Asset{}, err
	}
	asset := mapper.AssetToDomain(dbAsset)

	return asset, nil
}

func (r *PgContentRepository) DeleteAsset(ctx context.Context, id string) error {
	err := r.queries.DeleteAssetById(ctx, util.ToPgUUID(id))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete asset by id", "err", err)

		return err
	}

	return nil
}

func (r *PgContentRepository) GetAssets(ctx context.Context) ([]model.Asset, error) {
	dbAssets, err := r.queries.GetAllAssets(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get all assets", "err", err)

		return nil, err
	}

	assets := coreUtil.MapSlice(dbAssets, mapper.AssetToDomain)

	return assets, nil
}

func (r *PgContentRepository) GetTaskById(ctx context.Context, id string) (*model.Task, error) {
	dbTask, err := r.queries.GetTaskById(ctx, util.ToPgUUID(id))
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		slog.ErrorContext(ctx, "Failed to get task by id", "err", err)

		return nil, err
	}

	task, err := mapper.GetTaskByIdToDomain(dbTask)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to map task", "err", err)

		return nil, err
	}

	return &task, nil
}
