package postgres

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	specialtyEntity struct {
		ID     string
		Code   string
		Title  string
		UgsnID string
	}
)

func newSpecialty(entity specialtyEntity) *education.Speciality {
	specialty, _ := education.NewSpeciality(education.SpecialityParams{
		ID:     entity.ID,
		Code:   entity.Code,
		Title:  entity.Title,
		UgsnID: entity.UgsnID,
	})

	return specialty
}

func newSpecialties(entities []specialtyEntity) []*education.Speciality {
	specialties := make([]*education.Speciality, 0, len(entities))

	for _, document := range entities {
		specialties = append(specialties, newSpecialty(document))
	}

	return specialties
}

func newSpecialtyModels(entities []specialtyEntity) []query.SpecialtyModel {
	result := make([]query.SpecialtyModel, 0, len(entities))

	for _, document := range entities {
		result = append(result, query.SpecialtyModel{
			ID:     document.ID,
			Code:   document.Code,
			Title:  document.Title,
			UgsnID: document.UgsnID,
		})
	}

	return result
}
