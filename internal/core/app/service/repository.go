package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

var (
	ErrLevelNotFound          = errors.New("level not found")
	ErrUgsnNotFound           = errors.New("ugsn not found")
	ErrSpecialtyNotFound      = errors.New("specialty not found")
	ErrProgramAlreadyExists   = errors.New("program already exists")
	ErrSpecialtyAlreadyExists = errors.New("specialty already exists")
	ErrUgsnAlreadyExists      = errors.New("ugsn already exists")
	ErrLevelAlreadyExists     = errors.New("level already exists")
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
		AddUgsn(ctx context.Context, level *education.Level) error
		GetUgsn(ctx context.Context, id string) (*education.Level, error)
	}
)

type (
	SpecialtyRepository interface {
		AddSpecialty(ctx context.Context, level *education.Level) error
		GetSpecialty(ctx context.Context, id string) (*education.Level, error)
	}
)

type (
	ProgramRepository interface {
		AddProgram(ctx context.Context, level *education.Level) error
		GetProgram(ctx context.Context, id string) (*education.Level, error)
	}
)
