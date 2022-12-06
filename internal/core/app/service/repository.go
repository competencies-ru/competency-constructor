package service

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/pkg/errors"
)

var (
	ErrUgsnNotFound           = errors.New("ugsn: not found")
	ErrUgsnAlreadyExists      = errors.New("ugsn: already exists")
	ErrSpecialtyNotFound      = errors.New("specialty: not found")
	ErrSpecialtyAlreadyExists = errors.New("specialty: already exists")
	ErrProgramNotFound        = errors.New("program: already exists")
	ErrProgramAlreadyExists   = errors.New("program: already exists")
)

type (
	UgsnRepository interface {
		GetUgsn(ctx context.Context, code string) (*specialty.Ugsn, error)
		GetAllUgsn(ctx context.Context) ([]*specialty.Ugsn, error)
		AddUgsn(ctx context.Context, ugsn *specialty.Ugsn) error
		Exist(ctx context.Context, code string) bool
		FindUgsn(ctx context.Context, code string) (SpecificUgsn, error)
		UpdateUgsn(ctx context.Context, code string, u UgsnUpdater) error
	}

	UgsnUpdater func(
		ctx context.Context,
		ugns *specialty.Ugsn,
	) (*specialty.Ugsn, error)
)

type (
	SpecialtyRepository interface {
		GetAllSpecialty(ctx context.Context, ugsnCode string) []*specialty.Ugsn
		FindSpecialty(ctx context.Context, code string) (SpecificSpecialty, error)
		Exist(ctx context.Context, code string) bool
		UpdateSpecialty(ctx context.Context, code string, u SpecialtyUpdater) error
	}

	SpecialtyUpdater func(
		ctx context.Context,
		speciality *specialty.Speciality,
	) (*specialty.Speciality, error)
)
