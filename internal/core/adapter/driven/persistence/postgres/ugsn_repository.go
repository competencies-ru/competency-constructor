package postgres

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type UgsnRepository struct {
	db *pgxpool.Pool
}

func NewUgsnRepository(pool *pgxpool.Pool) *UgsnRepository {
	return &UgsnRepository{db: pool}
}

func (r *UgsnRepository) AddUgsn(
	ctx context.Context,
	ugsn *education.Ugsn,
) error {
	_, err := r.db.Exec(ctx, insertUgsn, ugsn.ID(), ugsn.Code(), ugsn.Title(), ugsn.LeveID())
	return err
}

func (r *UgsnRepository) GetUgsn(
	ctx context.Context,
	id string,
) (*education.Ugsn, error) {
	var entity = ugsnEntity{}
	if err := pgxscan.Get(ctx, r.db, entity, getUgsn, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.NotFoundEntity
		}
		return nil, err
	}
	return newUgsn(entity), nil
}

func (r *UgsnRepository) Update(ctx context.Context, id string, update service.UgsnUpdate) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	ugsn, err := r.GetUgsn(ctx, id)
	if err != nil {
		return err
	}

	ugsnUpdate, err := update(ctx, ugsn)
	if err != nil {
		return nil
	}

	if _, err = r.db.Exec(
		ctx,
		updateUgsn,
		ugsnUpdate.Code(),
		ugsnUpdate.Title(),
		ugsnUpdate.LeveID(),
		ugsnUpdate.ID(),
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *UgsnRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, ctx)

	res, err := tx.Exec(ctx, deleteUgsn, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *UgsnRepository) FindAllUgsn(ctx context.Context, levelID string) ([]query.UgsnModel, error) {
	var entities []ugsnEntity
	if err := pgxscan.Select(ctx, r.db, &entities, getUgsnsByLevelID, levelID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return newUgsnModels(entities), nil
}
