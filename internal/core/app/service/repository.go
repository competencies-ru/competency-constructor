package service

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/pkg/errors"
)

var (
	ErrUgsnNotFound           = errors.New("ugsn: not found")
	ErrUgsnAlreadyExists      = errors.New("ugsn: already exists")
	ErrSpecialtyNotFound      = errors.New("education: not found")
	ErrSpecialtyAlreadyExists = errors.New("education: already exists")
	ErrProgramNotFound        = errors.New("program: already exists")
	ErrProgramAlreadyExists   = errors.New("program: already exists")
)

type (
	UgsnRepository interface {
		GetUgsn(ctx context.Context, code string) (*education.Ugsn, error)
		GetAllUgsn(ctx context.Context) ([]*education.Ugsn, error)
		AddUgsn(ctx context.Context, ugsn *education.Ugsn) error
		Exist(ctx context.Context, code string) (bool, error)
		FindUgsn(ctx context.Context, code string) (*SpecificUgsn, error)
		UpdateUgsn(ctx context.Context, code string, u UgsnUpdater) error
	}

	UgsnUpdater func(
		ctx context.Context,
		ugns *education.Ugsn,
	) (*education.Ugsn, error)
)

type (
	SpecialtyRepository interface {
		GetAllSpecialty(ctx context.Context, ugsnCode string) []*education.Ugsn
		FindSpecialty(ctx context.Context, code string) (SpecificSpecialty, error)
		Exist(ctx context.Context, code string) bool
		UpdateSpecialty(ctx context.Context, code string, u SpecialtyUpdater) error
	}

	SpecialtyUpdater func(
		ctx context.Context,
		speciality *education.Speciality,
	) (*education.Speciality, error)
)
