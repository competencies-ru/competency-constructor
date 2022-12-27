package query

import (
	"context"
)

type SpecificLevelReadModel interface {
	FindLevel(ctx context.Context, id string) (SpecificLevelModel, error)
}

type SpecificLevelHandler interface {
	Handle(ctx context.Context, id string) (SpecificLevelModel, error)
}

type specificLevelHandler struct {
	readModel SpecificLevelReadModel
}

func NewSpecificLevelHandler(repo SpecificLevelReadModel) SpecificLevelHandler {
	if repo == nil {
		panic("repository is nil")
	}

	return specificLevelHandler{readModel: repo}
}

func (h specificLevelHandler) Handle(ctx context.Context, id string) (SpecificLevelModel, error) {
	return h.readModel.FindLevel(ctx, id)
}
