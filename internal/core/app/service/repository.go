package service

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"

	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

var (
	ErrLevelNotFound           = errors.New("level not found")
	ErrUgsnNotFound            = errors.New("ugsn not found")
	ErrSpecialtyNotFound       = errors.New("specialty not found")
	ErrProgramNotFound         = errors.New("program not found")
	ErrProgramAlreadyExists    = errors.New("program already exists")
	ErrCompetencyAlreadyExists = errors.New("competency already exists")
	ErrSpecialtyAlreadyExists  = errors.New("specialty already exists")
	ErrUgsnAlreadyExists       = errors.New("ugsn already exists")
	ErrLevelAlreadyExists      = errors.New("level already exists")
)

type (
	LevelRepository interface {
		AddLevel(ctx context.Context, level *education.Level) error
		GetLevel(ctx context.Context, id string) (*education.Level, error)
		UpdateLevel(ctx context.Context, id string, updater LevelUpdate) error
	}

	LevelUpdate func(ctx context.Context, level *education.Level) (*education.Level, error)
)

type (
	UgsnRepository interface {
		AddUgsn(ctx context.Context, ugsn *education.Ugsn) error
		GetUgsn(ctx context.Context, id string) (*education.Ugsn, error)
	}
)

type (
	SpecialtyRepository interface {
		AddSpecialty(ctx context.Context, level *education.Speciality) error
		GetSpecialty(ctx context.Context, id string) (*education.Speciality, error)
	}
)

type (
	ProgramRepository interface {
		AddProgram(ctx context.Context, level *education.Program) error
		GetProgram(ctx context.Context, id string) (*education.Program, error)
	}
)

type (
	CompetencyRepository interface {
		AddCompetency(ctx context.Context, competency *competencies.Competency) error
		GetCompetency(ctx context.Context, id string) (*competencies.Competency, error)
	}
)

type (
	IndicatorRepository interface {
		AddIndicator(ctx context.Context, indicator *competencies.Indicator) error
		GetIndicator(ctx context.Context, id string) (*competencies.Indicator, error)
		UpdateIndicator(ctx context.Context, id string, updater IndicatorUpdate) error
	}

	IndicatorUpdate func(ctx context.Context, level *competencies.Indicator) (*competencies.Indicator, error)
)

func IsNotFoundEntity(err error) bool {
	return errors.Is(err, ErrUgsnNotFound) ||
		errors.Is(err, ErrLevelNotFound) ||
		errors.Is(err, ErrSpecialtyNotFound) ||
		errors.Is(err, ErrProgramNotFound)
}
