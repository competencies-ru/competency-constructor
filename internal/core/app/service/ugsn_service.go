package service

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/multierr"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	UgsnCreate struct {
		Code  string
		Title string
	}

	UgsnUpdate struct {
		Title           string
		NewSpecialty    []Specialty
		DeleteSpecialty []string
	}
)

type UgsnHandler interface {
	Create(ctx context.Context, create UgsnCreate) error
	GetAllUgsn(ctx context.Context) ([]*education.Ugsn, error)
	GetSpecificUgsn(ctx context.Context, code string) (*SpecificUgsn, error)
	Update(ctx context.Context, code string, u UgsnUpdate) error
}

type ugsnHandler struct {
	ugsnRepo UgsnRepository
}

func NewUgsnHandler(ugsnRepo UgsnRepository) UgsnHandler {
	if ugsnRepo == nil {
		panic("ugsn repository is nil")
	}

	return ugsnHandler{
		ugsnRepo: ugsnRepo,
	}
}

func (u ugsnHandler) Create(ctx context.Context, create UgsnCreate) error {
	ugsn, err := education.NewUgsn(education.UgsnParams{
		Code:  create.Code,
		Title: create.Title,
	})
	if err != nil {
		return err
	}

	return u.ugsnRepo.AddUgsn(ctx, ugsn)
}

func (u ugsnHandler) GetAllUgsn(ctx context.Context) ([]*education.Ugsn, error) {
	return u.ugsnRepo.GetAllUgsn(ctx)
}

func (u ugsnHandler) GetSpecificUgsn(ctx context.Context, code string) (*SpecificUgsn, error) {
	if err := education.IsValidUgsnCode(code); err != nil {
		return nil, err
	}

	return u.ugsnRepo.FindUgsn(ctx, code)
}

func (u ugsnHandler) Update(ctx context.Context, code string, update UgsnUpdate) error {
	err := u.ugsnRepo.UpdateUgsn(ctx, code, updateUgsn(update))

	return errors.Wrapf(
		err,
		"update ugsn: %s", code)
}

func updateUgsn(u UgsnUpdate) UgsnUpdater {
	return func(ctx context.Context, ugns *education.Ugsn) (*education.Ugsn, error) {
		if err := ugns.Rename(u.Title); err != nil {
			return nil, err
		}

		var result error

		for _, v := range u.DeleteSpecialty {
			err := ugns.DeleteSpecialty(v)

			result = multierr.Append(result, err)
		}

		for _, s := range u.NewSpecialty {
			err := ugns.AddSpeciality(education.SpecialityParams{
				Code:     s.Code,
				Title:    s.Title,
				UgsnCode: ugns.Code(),
			})

			result = multierr.Append(result, err)
		}

		return ugns, result
	}
}
