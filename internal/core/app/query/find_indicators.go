package query

import (
	"context"
	"github.com/google/uuid"
)

type IndicatorReadModels interface {
	FindAllIndicators(
		ctx context.Context,
		competencyID string,
	) ([]IndicatorModel, error)
}

type FindAllIndicatorsHandler interface {
	Handle(ctx context.Context, competencyID uuid.UUID) ([]IndicatorModel, error)
}

type findAllIndicatorsHandler struct {
	repository IndicatorReadModels
}

func NewFindAllIndicatorsHandler(repository IndicatorReadModels) FindAllIndicatorsHandler {
	if repository == nil {
		panic("competency repository is nil ")
	}

	return findAllIndicatorsHandler{repository: repository}
}

func (h findAllIndicatorsHandler) Handle(ctx context.Context, competencyID uuid.UUID) ([]IndicatorModel, error) {
	return h.repository.FindAllIndicators(ctx, competencyID.String())
}
