package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// swagger:response answerPostResponse
type answerPostResponse struct {
	// in: body
	Body dto.AnswerResult
}

// swagger:parameters answerPost
type answerPostParams struct {
	// in: body
	Body dto.Answer
}

type answerHandler struct {
	validate   *validator.Validate
	saveAnswer *usecase.SaveAnswer
}

func NewAnswerHandler(validate *validator.Validate, saveAnswer *usecase.SaveAnswer) *answerHandler {
	return &answerHandler{validate: validate, saveAnswer: saveAnswer}
}

func (h *answerHandler) Register(router chi.Router) {
	router.Route("/v1/answer", func(r chi.Router) {
		// swagger:route POST /api/v1/answer answer answerPost
		//
		// Answer to the task
		//
		// Responses:
		//   200: answerPostResponse
		//   401: errorResponse
		//   404: errorResponse
		//   500: errorResponse
		r.Post("/", h.handlePostAnswer)
	})
}

func (h *answerHandler) handlePostAnswer(w http.ResponseWriter, r *http.Request) {
	body, err := util.DeserializeAndValidateBody[dto.Answer](r, h.validate)
	if err != nil {
		util.WriteError(w, err)
		return
	}
	result, err := h.saveAnswer.Invoke(r.Context(), dto.AnswerToModel(body))
	if err != nil {
		util.WriteError(w, err)
		return
	}

	util.WriteResponseJson(w, dto.AnswerResultFromModel(result.AnswerResult))
}
