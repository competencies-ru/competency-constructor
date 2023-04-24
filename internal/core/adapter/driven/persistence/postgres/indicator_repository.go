package postgres

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IndicatorEntity struct {
	ID           string
	Title        string
	Code         string
	SubjectID    string
	CompetencyID string
}

type IndicatorRepository struct {
	db *pgxpool.Pool
}

func NewIndicatorRepository(db *pgxpool.Pool) *IndicatorRepository {
	return &IndicatorRepository{
		db: db,
	}
}

func (r *IndicatorRepository) AddIndicator(ctx context.Context, ind *competencies.Indicator) error {
	_, err := r.db.Exec(ctx, insertIndicator, ind.ID(), ind.Title(), ind.Code(), ind.SubjectID(), ind.CompetencyID())
	return err
}

func (r *IndicatorRepository) Update(ctx context.Context, id string, updater service.IndicatorUpdate) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	indicator, err := r.GetIndicator(ctx, id)
	if err != nil {
		return err
	}

	upd, err := updater(ctx, indicator)
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(
		ctx,
		updateIndicator,
		upd.Title(),
		upd.Code(),
		upd.SubjectID(),
		upd.CompetencyID(),
		upd.ID(),
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *IndicatorRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	res, err := r.db.Exec(ctx, getIndicator, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *IndicatorRepository) GetIndicator(ctx context.Context, id string) (*competencies.Indicator, error) {
	var entity indicatorEntity
	err := pgxscan.Select(ctx, r.db, &entity, getIndicator, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, service.NotFoundEntity
		}

		return nil, err
	}

	return newIndicator(entity), err
}

func (r *IndicatorRepository) FindAllByCompetencyID(ctx context.Context, competencyID string) ([]query.IndicatorModel, error) {

	rows, err := r.db.Query(ctx, findAllIndicators, competencyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []indicatorEntity

	for rows.Next() {
		var indicator indicatorEntity
		var subject subjectEntity

		err := rows.Scan(
			&indicator.ID,
			&indicator.Title,
			&indicator.Code,
			&indicator.CompetencyID,
			&subject.ID,
			&subject.Name,
			&subject.SName,
		)
		if err != nil {
			return nil, err
		}

		indicator.Subject = subject

		entities = append(entities, indicator)
	}

	return newIndicatorModels(entities), err
}
