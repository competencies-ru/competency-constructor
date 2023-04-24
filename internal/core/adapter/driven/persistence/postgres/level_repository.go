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

type LevelRepository struct {
	db *pgxpool.Pool
}

func NewLevelRepository(pool *pgxpool.Pool) *LevelRepository {
	return &LevelRepository{db: pool}
}

func (r *LevelRepository) AddLevel(ctx context.Context, lvl *education.Level) error {
	_, err := r.db.Exec(ctx, insertLevel, lvl.ID(), lvl.Title())
	if err != nil {
		return err
	}

	return nil
}

func (r *LevelRepository) GetLevel(ctx context.Context, id string) (*education.Level, error) {
	var lvl levelEntity
	err := pgxscan.Get(ctx, r.db, &lvl, getLevel, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.NotFoundEntity
		}
		return nil, err
	}

	return newLevel(lvl), nil
}

func (r *LevelRepository) UpdateLevel(ctx context.Context, id string, updater service.LevelUpdate) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	level, err := r.GetLevel(ctx, id)
	if err != nil {
		return err
	}

	levelResult, err := updater(ctx, level)
	if err != nil {
		return err
	}

	res, err := r.db.Exec(ctx, updateLevel, levelResult.ID(), level.Title())
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *LevelRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	res, err := tx.Exec(ctx, deleteLevel, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return tx.Commit(ctx)
}

func (r *LevelRepository) FindLevels(ctx context.Context) ([]query.LevelModel, error) {
	var levels []levelEntity
	if err := pgxscan.Select(ctx, r.db, &levels, getLevels); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return newLevelModels(levels), nil
}
