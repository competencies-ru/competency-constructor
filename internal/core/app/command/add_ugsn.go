package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type CreateUgsnCommand struct {
	Code  string
	Title string
}

type AddUgsnHandler interface {
	Handle(ctx context.Context, levelID string, command []CreateUgsnCommand) error
}

type addUgsnHandler struct {
	repository service.LevelRepository
}

func NewAddUgsnHandler(repo service.LevelRepository) AddUgsnHandler {
	if repo == nil {
		panic("repository for add ugsn handler is nil")
	}

	return addUgsnHandler{repository: repo}
}

func (h addUgsnHandler) Handle(ctx context.Context, levelID string, commands []CreateUgsnCommand) error {
	err := h.repository.UpdateLevel(ctx, levelID, h.addUgsn(commands))

	return errors.Wrapf(err, "adding ugsn to level %s", levelID)
}

func (h addUgsnHandler) addUgsn(command []CreateUgsnCommand) service.LevelUpdate {
	return func(ctx context.Context, level *education.Level) (*education.Level, error) {
		var errResult error

		for _, ugsnCommand := range command {
			id := uuid.NewString()

			err := level.AddUgsn(education.UgsnParams{
				ID:    id,
				Code:  ugsnCommand.Code,
				Title: ugsnCommand.Title,
			})

			errResult = multierr.Append(errResult, err)
		}

		return level, errResult
	}
}
