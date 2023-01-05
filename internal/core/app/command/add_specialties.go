package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
)

type CreateSpecialtyCommand struct {
	Code  string
	Title string
}

type AddSpecialtiesHandler interface {
	Handle(ctx context.Context, uid string, command CreateSpecialtyCommand) error
}

type addSpecialtiesHandler struct {
}

func NewAddSpecialtiesHandler(repo service.LevelRepository) AddSpecialtiesHandler {
	if repo == nil {
		panic("repository for add specialties handler is nil")
	}

	return addSpecialtiesHandler{}
}

func (h addSpecialtiesHandler) Handle(ctx context.Context, uid string, commands CreateSpecialtyCommand) error {
	panic("")
}
