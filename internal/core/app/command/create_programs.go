package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateProgramCommand struct {
	Code  string
	Title string
}

type AddProgramHandler interface {
	Handle(ctx context.Context, sid string, commands CreateProgramCommand) error
}

type addProgramHandler struct {
	repository          service.ProgramRepository
	specialtyRepository service.SpecialtyRepository
}

func NewAddProgramsHandler(
	repo service.ProgramRepository,
	specialtyRepo service.SpecialtyRepository,
) AddProgramHandler {
	if repo == nil {
		panic("program repository is nil")
	}

	if specialtyRepo == nil {
		panic("specialty repository is nil")
	}

	return addProgramHandler{
		repository:          repo,
		specialtyRepository: specialtyRepo,
	}
}

func (h addProgramHandler) Handle(ctx context.Context, sid string, command CreateProgramCommand) (err error) {
	defer func() {
		err = errors.Wrapf(err, "find all programs by specialty code: %s", sid)
	}()

	specialty, err := h.specialtyRepository.GetSpecialty(ctx, sid)
	if err != nil {
		return err
	}

	if !education.MatchProgramCode(command.Code, specialty.Code()) {
		return education.ErrProgramNotMatchCode
	}

	pid := uuid.NewString()

	program, err := education.NewProgram(education.ProgramParams{
		ID:          pid,
		Code:        command.Code,
		Title:       command.Title,
		SpecialtyID: specialty.ID(),
	})
	if err != nil {
		return err
	}

	return h.repository.AddProgram(ctx, program)
}
