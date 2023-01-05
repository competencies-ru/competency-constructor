package query

import (
	"context"

	"github.com/pkg/errors"
)

type SpecialtiesReadModels interface {
	FindAllSpecialties(ctx context.Context, uid string) ([]SpecialtyModel, error)
}

type FindSpecialtiesHandler interface {
	Handle(ctx context.Context, ugsnCode string) ([]SpecialtyModel, error)
}

type findSpecialtiesHandler struct {
	levelRepo SpecialtiesReadModels
}

func NewFindSpecialtiesHandler(repo SpecialtiesReadModels) FindSpecialtiesHandler {
	if repo == nil {
		panic("level repository is nil")
	}

	return findSpecialtiesHandler{levelRepo: repo}
}

func (h findSpecialtiesHandler) Handle(ctx context.Context, uid string) ([]SpecialtyModel, error) {
	models, err := h.levelRepo.FindAllSpecialties(ctx, uid)

	return models, errors.Wrapf(err, "get ugsn by ugsn id %s", uid)
}
