package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type CreateProgramCommand struct {
	Code  string
	Title string
}

type AddProgramHandler interface {
	Handle(ctx context.Context, levelID, ugsnCode, scode string, commands []CreateProgramCommand) error
}

type addProgramHandler struct {
	repository service.LevelRepository
}

func NewAddProgramsHandler(repo service.LevelRepository) AddProgramHandler {
	if repo == nil {
		panic("repository for add specialties handler is nil")
	}

	return addProgramHandler{repository: repo}
}

func (h addProgramHandler) Handle(ctx context.Context, levelID, ugsnCode, scode string, commands []CreateProgramCommand) error {
	err := h.repository.UpdateLevel(ctx, levelID, h.addProgram(ugsnCode, scode, commands))

	return errors.Wrapf(err, "adding ugsn to level %s", levelID)
}

func (h addProgramHandler) addProgram(ugsCode, scode string, commands []CreateProgramCommand) service.LevelUpdate {
	return func(ctx context.Context, level *education.Level) (*education.Level, error) {
		var errResult error

		for _, command := range commands {
			id := uuid.NewString()

			err := level.AddProgram(ugsCode, scode, education.ProgramParams{
				ID:            id,
				Code:          command.Code,
				Title:         command.Title,
				SpecialtyCode: scode,
			})

			errResult = multierr.Append(errResult, err)
		}

		return level, errResult
	}
}
