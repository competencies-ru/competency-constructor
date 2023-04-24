package postgres

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type CompetencyRepository struct {
	competency *pgxpool.Pool
}

func NewCompetencyRepository(pool *pgxpool.Pool) *CompetencyRepository {
	return &CompetencyRepository{competency: pool}
}

func (r *CompetencyRepository) AddCompetency(
	ctx context.Context,
	competency *competencies.Competency,
) error {
	entity := newCompetencyEntity(competency)
	_, err := r.competency.Exec(
		ctx,
		insertCompetency,
		entity.ID,
		entity.Code,
		entity.Title,
		entity.Category,
		entity.CompetencyType,
		entity.LevelID,
		entity.UgsnID,
		entity.SpecialtyID,
		entity.ProgramID,
	)
	return err
}

func (r *CompetencyRepository) GetCompetency(
	ctx context.Context,
	id string,
) (*competencies.Competency, error) {
	var entity competencyEntity
	err := pgxscan.Select(ctx, r.competency, &entity, getCompetency, id)
	if err == pgx.ErrNoRows {
		return nil, service.NotFoundEntity
	}
	if err != nil {
		return nil, err
	}
	return newCompetency(entity), nil
}

func (r *CompetencyRepository) FilterCompetencies(
	ctx context.Context,
	levelID,
	ugsnID,
	specialtyID,
	programID string,
) ([]query.CompetencyModel, error) {
	var (
		nullLevelID     = stringToNullString(levelID)
		nullUgsnID      = stringToNullString(ugsnID)
		nullSpecialtyID = stringToNullString(specialtyID)
		nullProgramID   = stringToNullString(programID)
	)

	var entities []competencyEntity

	if err := pgxscan.Select(
		ctx,
		r.competency,
		&entities,
		filterCompetency,
		nullUgsnID,
		nullSpecialtyID,
		nullProgramID,
		nullLevelID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return newCompetencyModels(entities), nil
}

func (r *CompetencyRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.competency.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	res, err := r.competency.Exec(ctx, deleteCompetency, id)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *CompetencyRepository) Update(ctx context.Context, id string, updater service.CompetencyUpdate) error {
	tx, err := r.competency.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	competency, err := r.GetCompetency(ctx, id)
	if err != nil {
		return err
	}

	upd, err := updater(ctx, competency)
	if err != nil {
		return err
	}

	if _, err := r.competency.Exec(ctx, updateCompetency, upd.Code(), upd.Title(), upd.Category(), upd.ID()); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
