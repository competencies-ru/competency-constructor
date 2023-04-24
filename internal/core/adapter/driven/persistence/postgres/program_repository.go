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

type ProgramRepository struct {
	program *pgxpool.Pool
}

func NewProgramRepository(db *pgxpool.Pool) *ProgramRepository {
	return &ProgramRepository{program: db}
}

func (r *ProgramRepository) AddProgram(
	ctx context.Context, p *education.Program,
) error {
	_, err := r.program.Exec(ctx, insertProgram, p.ID(), p.Code(), p.Title(), p.SpecialtyID())
	return err
}

func (r *ProgramRepository) GetProgram(
	ctx context.Context,
	id string,
) (*education.Program, error) {
	var pEntity programEntity
	err := pgxscan.Select(ctx, r.program, &pEntity, getProgram, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, service.NotFoundEntity
		}
		return nil, err
	}
	return newProgram(pEntity), nil
}

func (r *ProgramRepository) Update(
	ctx context.Context,
	id string,
	updater service.ProgramUpdate,
) error {
	tx, err := r.program.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	program, err := r.GetProgram(ctx, id)
	if err != nil {
		return err
	}

	program, err = updater(ctx, program)

	if err != nil {
		return err
	}

	if _, err := r.program.Exec(ctx, updateProgram, program.Code(), program.Title(), program.ID()); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *ProgramRepository) Delete(
	ctx context.Context,
	id string,
) error {
	tx, err := r.program.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	res, err := tx.Exec(ctx, deleteProgram, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *ProgramRepository) FindAllPrograms(
	ctx context.Context,
	specialtiesID string,
) ([]query.ProgramModel, error) {
	var p []programEntity
	err := pgxscan.Select(ctx, r.program, &p, getProgramBySpecialtyID, specialtiesID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []query.ProgramModel{}, nil
		}
		return nil, err
	}
	return newProgramModels(p), nil
}
