package service

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"

	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

var (
	NotFoundEntity      = errors.New("not found")
	AlreadyExistsEntity = errors.New("already exist")
)

type (
	LevelRepository interface {
		AddLevel(ctx context.Context, level *education.Level) error
		GetLevel(ctx context.Context, id string) (*education.Level, error)
		UpdateLevel(ctx context.Context, id string, updater LevelUpdate) error
		Delete(ctx context.Context, id string) error
	}

	LevelUpdate func(ctx context.Context, level *education.Level) (*education.Level, error)
)

type (
	UgsnRepository interface {
		AddUgsn(ctx context.Context, ugsn *education.Ugsn) error
		GetUgsn(ctx context.Context, id string) (*education.Ugsn, error)
		Update(ctx context.Context, id string, update UgsnUpdate) error
		Delete(ctx context.Context, id string) error
	}

	UgsnUpdate func(ctx context.Context, ugsn *education.Ugsn) (*education.Ugsn, error)
)

type (
	SpecialtyRepository interface {
		AddSpecialty(ctx context.Context, level *education.Speciality) error
		GetSpecialty(ctx context.Context, id string) (*education.Speciality, error)
		Update(ctx context.Context, id string, update SpecialtyUpdate) error
		Delete(ctx context.Context, id string) error
	}

	SpecialtyUpdate func(ctx context.Context, specialty *education.Speciality) (*education.Speciality, error)
)

type (
	ProgramRepository interface {
		AddProgram(ctx context.Context, program *education.Program) error
		GetProgram(ctx context.Context, id string) (*education.Program, error)
		Update(ctx context.Context, id string, update ProgramUpdate) error
		Delete(ctx context.Context, id string) error
	}

	ProgramUpdate func(ctx context.Context, program *education.Program) (*education.Program, error)
)

type (
	CompetencyRepository interface {
		AddCompetency(ctx context.Context, competency *competencies.Competency) error
		GetCompetency(ctx context.Context, id string) (*competencies.Competency, error)
		Update(ctx context.Context, id string, update CompetencyUpdate) error
		Delete(ctx context.Context, id string) error
	}

	CompetencyUpdate func(ctx context.Context, competency *competencies.Competency) (*competencies.Competency, error)
)

type (
	IndicatorRepository interface {
		AddIndicator(ctx context.Context, indicator *competencies.Indicator) error
		GetIndicator(ctx context.Context, id string) (*competencies.Indicator, error)
		Update(ctx context.Context, id string, updater IndicatorUpdate) error
		Delete(ctx context.Context, id string) error
	}

	IndicatorUpdate func(ctx context.Context, level *competencies.Indicator) (*competencies.Indicator, error)
)

func IsNotFoundEntity(err error) bool {
	return errors.Is(err, NotFoundEntity)
}

func IsAlreadyExistEntity(err error) bool {
	return errors.Is(err, AlreadyExistsEntity)
}
