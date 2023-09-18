package port

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

type AnswerRepository interface {
	SaveHistoricalAnswer(context.Context, model.HistoricalAnswer) error
}
