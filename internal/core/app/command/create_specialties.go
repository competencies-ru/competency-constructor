package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
)

type CreateSpecialtyCommand struct {
	Code  string
	Title string
}

type AddSpecialtiesHandler interface {
	Handle(ctx context.Context, uid string, command CreateSpecialtyCommand) error
}

type addSpecialtiesHandler struct {
	ugsnRepository service.UgsnRepository
	repository     service.SpecialtyRepository
}

func NewAddSpecialtiesHandler(
	repo service.SpecialtyRepository,
	ugsnRepo service.UgsnRepository,
) AddSpecialtiesHandler {
	if repo == nil {
		panic("service repository is nil")
	}

	if ugsnRepo == nil {
		panic("ugsn repository is nil")
	}

	return addSpecialtiesHandler{
		repository:     repo,
		ugsnRepository: ugsnRepo,
	}
}

func (h addSpecialtiesHandler) Handle(ctx context.Context, uid string, command CreateSpecialtyCommand) (err error) {
	defer func() {
		err = errors.Wrapf(err, "find all specialties by ugsn id: %s", uid)
	}()

	ugsn, err := h.ugsnRepository.GetUgsn(ctx, uid)
	if err != nil {
		return err
	}

	if !education.MatchSpecialtyCode(command.Code, ugsn.Code()) {
		return education.ErrSpecialityNotMatchCode
	}

	pid := uuid.NewString()

	specialty, err := education.NewSpeciality(education.SpecialityParams{
		ID:     pid,
		Code:   command.Code,
		Title:  command.Title,
		UgsnID: ugsn.ID(),
	})
	if err != nil {
		return err
	}

	return h.repository.AddSpecialty(ctx, specialty)
}
