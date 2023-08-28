package mapper

import (
	"encoding/json"
	"log/slog"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

func TechnologyToDomain(technology db.Technology) model.Technology {
	return model.Technology{
		Id:          util.FromPgUUID(technology.ID),
		Title:       technology.Title,
		Description: util.FromPgString(technology.Description),
		ImageUrl:    util.FromPgString(technology.ImageUrl),
		Position:    int(technology.Position),
	}
}

func SectionToDomain(section db.Section) model.Section {
	return model.Section{
		Id:           util.FromPgUUID(section.ID),
		TechnologyId: util.FromPgUUID(section.TechnologyID),
		Title:        section.Title,
		Description:  util.FromPgString(section.Description),
		ImageUrl:     util.FromPgString(section.ImageUrl),
		Position:     int(section.Position),
	}
}

func TaskToDomain(task db.Task) (model.Task, error) {
	var content model.TaskContent
	err := json.Unmarshal(task.Content, &content)
	if err != nil {
		slog.Error("Failed to unmarshal task content", "err", err)
		return model.Task{}, err
	}

	return model.Task{
		Id:          util.FromPgUUID(task.ID),
		SectionId:   util.FromPgUUID(task.SectionID),
		Title:       task.Title,
		Description: util.FromPgString(task.Description),
		Position:    util.FromPgInt(task.Position),
		ImageUrl:    util.FromPgString(task.ImageUrl),
		Difficulty:  int(task.Difficulty),
		IsPublic:    task.IsPublic,
		Content:     content,
	}, nil
}

func UpsertTechnologyParamsFromDomain(technology model.Technology) db.UpsertTechnologyParams {
	return db.UpsertTechnologyParams{
		ID:          util.ToPgUUID(technology.Id),
		Title:       technology.Title,
		Description: util.ToPgString(technology.Description),
		ImageUrl:    util.ToPgString(technology.ImageUrl),
		Position:    int32(technology.Position),
	}
}

func UpsertSectionParamsFromDomain(section model.Section) db.UpsertSectionParams {
	return db.UpsertSectionParams{
		ID:           util.ToPgUUID(section.Id),
		TechnologyID: util.ToPgUUID(section.TechnologyId),
		Title:        section.Title,
		Description:  util.ToPgString(section.Description),
		ImageUrl:     util.ToPgString(section.ImageUrl),
		Position:     int32(section.Position),
	}
}

func UpsertTaskParamsFromDomain(task model.Task) (db.UpsertTaskParams, error) {
	content, err := json.Marshal(task.Content)
	if err != nil {
		return db.UpsertTaskParams{}, err
	}

	return db.UpsertTaskParams{
		ID:          util.ToPgUUID(task.Id),
		SectionID:   util.ToPgUUID(task.SectionId),
		Title:       task.Title,
		Description: util.ToPgString(task.Description),
		Position:    util.ToPgInt4(task.Position),
		ImageUrl:    util.ToPgString(task.ImageUrl),
		Difficulty:  int32(task.Difficulty),
		IsPublic:    task.IsPublic,
		Content:     content,
	}, nil
}

func GetAllTechnologiesWithSectionsPreviewRowsToDomain(
	rows []db.GetAllTechnologiesWithSectionsPreviewRow,
) []model.TechnologyWithSectionsPreview {
	return util.MapJoinedRows[db.GetAllTechnologiesWithSectionsPreviewRow, model.TechnologyWithSectionsPreview, model.SectionPreview](
		rows,
		func(row db.GetAllTechnologiesWithSectionsPreviewRow, sections []model.SectionPreview) model.TechnologyWithSectionsPreview {
			return model.TechnologyWithSectionsPreview{
				Technology: model.Technology{
					Id:          util.FromPgUUID(row.ID),
					Title:       row.Title,
					Description: util.FromPgString(row.Description),
					ImageUrl:    util.FromPgString(row.ImageUrl),
					Position:    int(row.Position),
				},
				Sections: sections,
			}
		},
		func(row db.GetAllTechnologiesWithSectionsPreviewRow) (parentID string, child model.SectionPreview) {
			return util.FromPgUUID(row.ID), model.SectionPreview{
				Id:    util.FromPgUUID(row.SectionID),
				Title: row.SectionTitle,
			}
		},
	)
}

func GetAllTechnolotySectionsWithTasksPreviewRowsToDomain(
	rows []db.GetAllTechnolotySectionsWithTasksPreviewRow,
) []model.SectionWithTasksPreview {
	return util.MapJoinedRows[db.GetAllTechnolotySectionsWithTasksPreviewRow, model.SectionWithTasksPreview, model.TaskPreview](
		rows,
		func(row db.GetAllTechnolotySectionsWithTasksPreviewRow, tasks []model.TaskPreview) model.SectionWithTasksPreview {
			return model.SectionWithTasksPreview{
				Section: model.Section{
					Id:           util.FromPgUUID(row.ID),
					TechnologyId: util.FromPgUUID(row.TechnologyID),
					Title:        row.Title,
					Description:  util.FromPgString(row.Description),
					ImageUrl:     util.FromPgString(row.ImageUrl),
				},
				Tasks: tasks,
			}
		},
		func(row db.GetAllTechnolotySectionsWithTasksPreviewRow) (parentID string, child model.TaskPreview) {
			return util.FromPgUUID(row.ID), model.TaskPreview{
				Id:       util.FromPgUUID(row.TaskID),
				Title:    row.TaskTitle,
				IsPublic: row.TaskIsPublic,
			}
		},
	)
}
