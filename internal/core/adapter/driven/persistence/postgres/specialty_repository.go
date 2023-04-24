package postgres

import (
	"context"
	"database/sql"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpecialtyRepository struct {
	db *pgxpool.Pool
}

func NewSpecialtyRepository(db *pgxpool.Pool) *SpecialtyRepository {
	return &SpecialtyRepository{db: db}
}

func (r *SpecialtyRepository) AddSpecialty(ctx context.Context, s *education.Speciality) error {
	_, err := r.db.Exec(ctx, insertSpecialty, s.ID(), s.Code(), s.Title(), s.UgsnID())
	return err
}

func (r *SpecialtyRepository) GetSpecialty(ctx context.Context, id string) (*education.Speciality, error) {
	var specialty = specialtyEntity{}
	err := pgxscan.Select(ctx, r.db, &specialty, getSpecialty, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, service.NotFoundEntity
		}
		return nil, err
	}
	return newSpecialty(specialty), nil
}

func (r *SpecialtyRepository) Update(ctx context.Context, id string, update service.SpecialtyUpdate) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	specialty, err := r.GetSpecialty(ctx, id)
	if err != nil {
		return err
	}

	s, err := update(ctx, specialty)
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(ctx, updateSpecialty, s.Code, s.Title, s.UgsnID, s.ID); err != nil {
		return nil
	}

	return tx.Commit(ctx)
}

func (r *SpecialtyRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	res, err := tx.Exec(ctx, deleteSpecialty, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *SpecialtyRepository) FindAllSpecialties(ctx context.Context, ugsnID string) ([]query.SpecialtyModel, error) {
	var specialties []specialtyEntity
	err := pgxscan.Select(ctx, r.db, &specialties, getSpecialtyByUgsnID, ugsnID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []query.SpecialtyModel{}, nil
		}
		return nil, err
	}
	return newSpecialtyModels(specialties), nil
}
