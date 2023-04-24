package postgres

import (
	"database/sql"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

type subjectEntity struct {
	ID    sql.NullString
	Name  sql.NullString
	SName sql.NullString
}

func newSubjectModel(entity subjectEntity) query.SubjectModel {
	return query.SubjectModel{
		ID:    entity.ID.String,
		Name:  entity.Name.String,
		SName: entity.SName.String,
	}
}

func newSubjectModels(documents []subjectEntity) []query.SubjectModel {
	buff := make([]query.SubjectModel, 0, len(documents))

	for _, document := range documents {
		buff = append(buff, newSubjectModel(document))
	}

	return buff
}

func newSubjectEntity(subject *competencies.Subject) subjectEntity {
	return subjectEntity{
		ID:    stringToNullString(subject.ID()),
		Name:  stringToNullString(subject.Name()),
		SName: stringToNullString(subject.SName()),
	}
}

func subjectIsEmpty(entity subjectEntity) bool {
	return entity.ID.Valid
}
