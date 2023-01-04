package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

var (
	ErrLevelNotFound       = errors.New("level not found")
	ErrUgsnNotFound        = errors.New("ugsn not found")
	ErrSpecialtiesNotFound = errors.New("specialties not found")
)

type (
	LevelRepository interface {
		AddLevel(ctx context.Context, level *education.Level) error
		GetLevel(ctx context.Context, id string) (*education.Level, error)
		UpdateLevel(ctx context.Context, id string, updater LevelUpdate) error
	}

	LevelUpdate func(ctx context.Context, level *education.Level) (*education.Level, error)
)

func IsNotFoundLevelEntities(err error) bool {
	return errors.Is(err, ErrLevelNotFound) ||
		errors.Is(err, ErrUgsnNotFound) ||
		errors.Is(err, ErrSpecialtiesNotFound)
}
