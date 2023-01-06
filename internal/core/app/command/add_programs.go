package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
)

type CreateProgramCommand struct {
	Code  string
	Title string
}

type AddProgramHandler interface {
	Handle(ctx context.Context, sid string, commands CreateProgramCommand) error
}

type addProgramHandler struct{}

func NewAddProgramsHandler(repo service.LevelRepository) AddProgramHandler {
	if repo == nil {
		panic("repository for add specialties handler is nil")
	}

	return addProgramHandler{}
}

func (h addProgramHandler) Handle(ctx context.Context, sid string, command CreateProgramCommand) error {
	panic("")
}
