package postgres

import (
	"context"
	"database/sql"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
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

func (u UgsnRepository) AddUgsn(
	ctx context.Context,
	ugsn *specialty.Ugsn,
) error {
	exist, err := u.Exist(ctx, ugsn.Code())
	if err != nil {
		return err
	}

	if exist {
		return errors.Wrapf(service.ErrUgsnAlreadyExists, "create ugsn by code: %s", ugsn.Code())
	}

	exec, err := u.db.Exec(ctx, createUgsn, ugsn.Code(), ugsn.Title())
	if err != nil {
		return err
	}

	exec.Insert()

	return nil
}

func (u UgsnRepository) Exist(ctx context.Context, code string) (bool, error) {
	var ok bool

	err := u.db.QueryRow(ctx, existUgsn, code).Scan(&ok)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (u UgsnRepository) FindUgsn(ctx context.Context, code string) (*service.SpecificUgsn, error) {
	su := &specificUgsn{Specialties: make(map[string]*specificSpecialty)}

	query, err := u.db.Query(ctx, findUgns, code)
	if err != nil {
		return nil, err
	}

	isFirst := true

	for query.Next() {
		var (
			scode      sql.NullString
			stitle     sql.NullString
			sugsn      sql.NullString
			pid        sql.NullString
			ptitle     sql.NullString
			pspecialty sql.NullString
		)

		if isFirst {
			err := query.Scan(&su.Code, &su.Title, &scode, &stitle, &sugsn, &pid, &ptitle, &pspecialty)
			if err != nil {
				return nil, err
			}
			isFirst = false
			unmarshallingSpecificUgsn(su, scode, stitle, sugsn, pid, ptitle, pspecialty)

			continue
		}

		err := query.Scan(nil, nil, &scode, &stitle, &sugsn, &pid, &ptitle, &pspecialty)
		if err != nil {
			return nil, err
		}

		unmarshallingSpecificUgsn(su, scode, stitle, sugsn, pid, ptitle, pspecialty)

	}

	return marshallingSpecificUgsn(su), err
}

func (u UgsnRepository) UpdateUgsn(ctx context.Context, code string, upt service.UgsnUpdater) error {
	// TODO implement me
	panic("implement me")
}
