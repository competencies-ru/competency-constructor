package query

import (
	"context"
)

type LevelReadModels interface {
	FindLevels(ctx context.Context) ([]LevelModel, error)
}

type FindLevelsHandler interface {
	Handle(ctx context.Context) ([]LevelModel, error)
}

type findLevelsHandler struct {
	levelRepo LevelReadModels
}

func NewFindLevelsHandler(repo LevelReadModels) FindLevelsHandler {
	if repo == nil {
		panic("level repository is nil")
	}

	return findLevelsHandler{levelRepo: repo}
}

func (h findLevelsHandler) Handle(ctx context.Context) ([]LevelModel, error) {
	return h.levelRepo.FindLevels(ctx)
}
