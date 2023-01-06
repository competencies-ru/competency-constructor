package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	specialtiesDocument struct {
		ID     string `bson:"_id,omitempty"`
		Code   string `bson:"code"`
		Title  string `bson:"title"`
		UgsnID string `bson:"ugsn_id"`
	}
)

func newSpecialtyDocument(specialty *education.Speciality) specialtiesDocument {
	return specialtiesDocument{
		ID:     specialty.ID(),
		Code:   specialty.Code(),
		Title:  specialty.Title(),
		UgsnID: specialty.UgsnID(),
	}
}

func newSpecialty(document specialtiesDocument) *education.Speciality {
	specialty, _ := education.NewSpeciality(education.SpecialityParams{
		ID:     document.ID,
		Code:   document.Code,
		Title:  document.Title,
		UgsnID: document.UgsnID,
	})

	return specialty
}

func newSpecialties(documents []specialtiesDocument) []*education.Speciality {
	specialties := make([]*education.Speciality, 0, len(documents))

	for _, document := range documents {
		specialties = append(specialties, newSpecialty(document))
	}

	return specialties
}

func newSpecialtyModels(documents []specialtiesDocument) []query.SpecialtyModel {
	result := make([]query.SpecialtyModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.SpecialtyModel{
			ID:     document.ID,
			Code:   document.Code,
			Title:  document.Title,
			UgsnID: document.UgsnID,
		})
	}

	return result
}
