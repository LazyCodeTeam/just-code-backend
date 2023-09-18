package mapper

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

func InsertAnswerParamsFromModel(answer model.HistoricalAnswer) db.InsertAnswerParams {
	return db.InsertAnswerParams{
		TaskID:    util.ToPgUUID(answer.TaskId),
		ProfileID: answer.ProfileId,
		Result:    db.AnswerResult(answer.AnswerResult),
	}
}
