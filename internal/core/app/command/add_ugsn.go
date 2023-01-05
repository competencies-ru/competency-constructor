package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
)

type CreateUgsnCommand struct {
	Code  string
	Title string
}

type AddUgsnHandler interface {
	Handle(ctx context.Context, levelID string, command CreateUgsnCommand) error
}

type addUgsnHandler struct {
}

func NewAddUgsnHandler(repo service.LevelRepository) AddUgsnHandler {
	if repo == nil {
		panic("repository for add ugsn handler is nil")
	}

	return addUgsnHandler{}
}

func (h addUgsnHandler) Handle(ctx context.Context, levelID string, command CreateUgsnCommand) error {
	panic("")
}
