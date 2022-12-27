package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type CreateSpecialtiesCommand struct {
	Code  string
	Title string
}

type AddSpecialtiesHandler interface {
	Handle(ctx context.Context, levelID, ugsnCode string, command []CreateSpecialtiesCommand) error
}

type addSpecialtiesHandler struct {
	repository service.LevelRepository
}

func NewAddSpecialtiesHandler(repo service.LevelRepository) AddSpecialtiesHandler {
	if repo == nil {
		panic("repository for add specialties handler is nil")
	}

	return addSpecialtiesHandler{repository: repo}
}

func (h addSpecialtiesHandler) Handle(ctx context.Context, levelID, ugsnCode string, commands []CreateSpecialtiesCommand) error {
	err := h.repository.UpdateLevel(ctx, levelID, h.addSpecialties(ugsnCode, commands))

	return errors.Wrapf(err, "adding ugsn to level %s", levelID)
}

func (h addSpecialtiesHandler) addSpecialties(ugsCode string, commands []CreateSpecialtiesCommand) service.LevelUpdate {
	return func(ctx context.Context, level *education.Level) (*education.Level, error) {
		var errResult error

		for _, command := range commands {
			id := uuid.NewString()

			err := level.AddSpecialty(ugsCode, education.SpecialityParams{
				ID:       id,
				Code:     command.Code,
				Title:    command.Title,
				UgsnCode: ugsCode,
			})

			errResult = multierr.Append(errResult, err)
		}

		return level, errResult
	}
}
