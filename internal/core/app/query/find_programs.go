package query

import (
	"context"

	"github.com/pkg/errors"
)

type ProgramsReadModels interface {
	FindAllPrograms(ctx context.Context, sui string) ([]ProgramModel, error)
}

type FindProgramsHandler interface {
	Handle(ctx context.Context, sid string) ([]ProgramModel, error)
}

type findProgramsHandler struct {
	levelRepo ProgramsReadModels
}

func NewFindProgramsHandler(repo ProgramsReadModels) FindProgramsHandler {
	if repo == nil {
		panic("level repository is nil")
	}

	return findProgramsHandler{levelRepo: repo}
}

func (h findProgramsHandler) Handle(ctx context.Context, sid string) ([]ProgramModel, error) {
	models, err := h.levelRepo.FindAllPrograms(ctx, sid)

	return models, errors.Wrapf(err, "get programs by specialty id: %s", sid)
}
