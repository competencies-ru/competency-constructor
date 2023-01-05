package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
)

type (
	CreateLevel struct {
		Title string
	}
)

type CreateLevelHandler interface {
	Handle(ctx context.Context, cmd CreateLevel) (string, error)
}

type createLevelHandler struct {
	levelRepo service.LevelRepository
}

func NewCreateLevelHandler(repo service.LevelRepository) CreateLevelHandler {
	if repo == nil {
		panic("level repository is nil")
	}

	return createLevelHandler{levelRepo: repo}
}

func (h createLevelHandler) Handle(ctx context.Context, cmd CreateLevel) (string, error) {
	id := uuid.NewString()

	level, err := education.NewLevel(education.LevelParam{
		ID:    id,
		Title: cmd.Title,
	})
	if err != nil {
		return "", err
	}

	return id, h.levelRepo.AddLevel(ctx, level)
}
