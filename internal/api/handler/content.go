package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type contentHandler struct {
	validate *validator.Validate
}

func NewContentHandler(validate *validator.Validate) *contentHandler {
	return &contentHandler{
		validate: validate,
	}
}

func (h *contentHandler) Register(router chi.Router) {
}
