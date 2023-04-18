package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

type subjectDocument struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	SName string `bson:"sname"`
}

func newSubjectModel(document subjectDocument) query.SubjectModel {
	return query.SubjectModel{
		ID:    document.ID,
		Name:  document.Name,
		SName: document.SName,
	}
}

func newSubjectModels(documents []subjectDocument) []query.SubjectModel {
	buff := make([]query.SubjectModel, 0, len(documents))

	for _, document := range documents {
		buff = append(buff, newSubjectModel(document))
	}

	return buff
}

func newSubjectDocument(subject *competencies.Subject) subjectDocument {
	return subjectDocument{
		ID:    subject.ID(),
		Name:  subject.Name(),
		SName: subject.SName(),
	}
}

func newSubjectDocuments(subjects []*competencies.Subject) []interface{} {
	buff := make([]interface{}, 0, len(subjects))

	for i := range subjects {
		tmp := subjects[i]
		buff = append(buff, newSubjectDocument(tmp))
	}

	return buff
}
