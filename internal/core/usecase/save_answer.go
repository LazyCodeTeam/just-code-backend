package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type SaveAnswer struct {
	transactionFactory port.TransactionFactory
}

func NewSaveAnswer(transactionFactory port.TransactionFactory) *SaveAnswer {
	return &SaveAnswer{transactionFactory: transactionFactory}
}

func (s *SaveAnswer) Invoke(
	ctx context.Context,
	answer model.Answer,
) (model.HistoricalAnswer, error) {
	transaction, err := s.transactionFactory.Begin(ctx)
	if err != nil {
		return model.HistoricalAnswer{}, err
	}
	defer transaction.Rollback(ctx)

	contentRepository := transaction.ContentRepository(ctx)
	answerRepository := transaction.AnswerRepository(ctx)
	task, err := contentRepository.GetTaskById(ctx, answer.TaskId)
	if err != nil {
		return model.HistoricalAnswer{}, err
	}

	if task == nil {
		return model.HistoricalAnswer{}, failure.NewNotFoundFailure(failure.FailureTypeNotFound)
	}

	result, err := task.IsAnswerValid(answer)
	if err != nil {
		return model.HistoricalAnswer{}, err
	}

	historicalAnswer := model.HistoricalAnswer{
		ProfileId:    *util.ExtractCurrentUserId(ctx),
		TaskId:       answer.TaskId,
		AnswerResult: result,
	}
	err = answerRepository.SaveHistoricalAnswer(ctx, historicalAnswer)
	if err != nil {
		return model.HistoricalAnswer{}, err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return model.HistoricalAnswer{}, err
	}

	return historicalAnswer, nil
}
