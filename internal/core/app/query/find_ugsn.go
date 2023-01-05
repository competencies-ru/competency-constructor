package query

import (
	"context"

	"github.com/pkg/errors"
)

type UgsnReadModels interface {
	FindAllUgsn(ctx context.Context, levelID string) ([]UgsnModel, error)
}

type FindUgsnHandler interface {
	Handle(ctx context.Context, levelID string) ([]UgsnModel, error)
}

type findUgsnHandler struct {
	levelRepo UgsnReadModels
}

func NewFindUgsnHandler(repo UgsnReadModels) FindUgsnHandler {
	if repo == nil {
		panic("level repository is nil")
	}

	return findUgsnHandler{levelRepo: repo}
}

func (h findUgsnHandler) Handle(ctx context.Context, levelID string) ([]UgsnModel, error) {
	models, err := h.levelRepo.FindAllUgsn(ctx, levelID)

	return models, errors.Wrapf(err, "get ugsn by level id: %s", levelID)
}
