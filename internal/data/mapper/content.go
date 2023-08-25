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
		Id:          util.DecodeUUID(technology.ID),
		Title:       technology.Title,
		Description: util.FromPgString(technology.Description),
		ImageUrl:    util.FromPgString(technology.ImageUrl),
		Position:    int(technology.Position),
	}
}

func SectionToDomain(section db.Section) model.Section {
	return model.Section{
		Id:          util.DecodeUUID(section.ID),
		Title:       section.Title,
		Description: util.FromPgString(section.Description),
		ImageUrl:    util.FromPgString(section.ImageUrl),
		Position:    int(section.Position),
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
		Id:          util.DecodeUUID(task.ID),
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
		ID:          util.EncodeUUID(technology.Id),
		Title:       technology.Title,
		Description: util.ToPgString(technology.Description),
		ImageUrl:    util.ToPgString(technology.ImageUrl),
		Position:    int32(technology.Position),
	}
}

func UpsertSectionParamsFromDomain(section model.Section) db.UpsertSectionParams {
	return db.UpsertSectionParams{
		ID:           util.EncodeUUID(section.Id),
		TechnologyID: util.EncodeUUID(section.TechnologyId),
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
		ID:          util.EncodeUUID(task.Id),
		SectionID:   util.EncodeUUID(task.SectionId),
		Title:       task.Title,
		Description: util.ToPgString(task.Description),
		Position:    util.ToPgInt4(task.Position),
		ImageUrl:    util.ToPgString(task.ImageUrl),
		Difficulty:  int32(task.Difficulty),
		IsPublic:    task.IsPublic,
		Content:     content,
	}, nil
}