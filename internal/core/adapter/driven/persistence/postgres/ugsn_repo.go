package postgres

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UgsnRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) UgsnRepository {
	if db == nil {
		panic("db connection is nil")
	}

	return UgsnRepository{db: db}
}

func (u UgsnRepository) GetUgsn(ctx context.Context, code string) (*specialty.Ugsn, error) {
	panic("implement me")
}

func (u UgsnRepository) GetAllUgsn(ctx context.Context) ([]*specialty.Ugsn, error) {
	var res []*ugsn

	err := pgxscan.Select(ctx, u.db, &res, getAllUgsn)
	if err != nil {
		return nil, err
	}

	return mapUgsn(res)
}

func (u UgsnRepository) AddUgsn(ctx context.Context, ugsn *specialty.Ugsn) error {
	// TODO implement me
	panic("implement me")
}

func (u UgsnRepository) Exist(ctx context.Context, code string) bool {
	// TODO implement me
	panic("implement me")
}

func (u UgsnRepository) FindUgsn(ctx context.Context, code string) (service.SpecificUgsn, error) {
	// TODO implement me
	panic("implement me")
}

func (u UgsnRepository) UpdateUgsn(ctx context.Context, code string, upt service.UgsnUpdater) error {
	// TODO implement me
	panic("implement me")
}
