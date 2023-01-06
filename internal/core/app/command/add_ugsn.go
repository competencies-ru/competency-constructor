package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"

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
	repository      service.UgsnRepository
	levelRepository service.LevelRepository
}

func NewAddUgsnHandler(
	repo service.UgsnRepository,
	levelRepo service.LevelRepository,
) AddUgsnHandler {
	if repo == nil {
		panic("ugsn repository is nil")
	}

	if levelRepo == nil {
		panic("level repository is nil")
	}

	return addUgsnHandler{repository: repo, levelRepository: levelRepo}
}

func (h addUgsnHandler) Handle(ctx context.Context, levelID string, command CreateUgsnCommand) (err error) {
	defer func() {
		err = errors.Wrapf(err, "find all ugsn by level id: %s", levelID)
	}()

	_, err = h.levelRepository.GetLevel(ctx, levelID)

	if err != nil {
		return err
	}

	uid := uuid.NewString()

	ugsn, err := education.NewUgsn(education.UgsnParams{
		ID:      uid,
		Code:    command.Code,
		Title:   command.Title,
		LevelID: levelID,
	})
	if err != nil {
		return err
	}

	return h.repository.AddUgsn(ctx, ugsn)
}
