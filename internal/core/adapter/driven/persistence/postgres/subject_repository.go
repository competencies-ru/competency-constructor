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

type SubjectRepository struct {
	db *pgxpool.Pool
}

func NewSubjectRepo(db *pgxpool.Pool) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func (r *SubjectRepository) AddSubject(ctx context.Context, entity *subjectEntity) error {
	_, err := r.db.Exec(ctx, insertSubject, entity.ID, entity.Name, entity.SName)
	return err
}

func (r *SubjectRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.Exec(ctx, deleteSubject, id)
	if res.RowsAffected() == 0 {
		return service.NotFoundEntity
	}

	return err
}

func (r *SubjectRepository) Update(ctx context.Context, subject *competencies.Subject) error {
	_, err := r.db.Exec(ctx, updateSubject, subject.ID(), subject.Name(), subject.SName())
	return err
}

func (r *SubjectRepository) FilterSubjectsByName(ctx context.Context, name string) ([]query.SubjectModel, error) {
	var entities []subjectEntity
	err := pgxscan.Select(ctx, r.db, &entities, filterSubjectByName, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	return newSubjectModels(entities), nil
}
