package mongodb

import "github.com/competencies-ru/competency-constructor/internal/core/entity/education"

type (
	specialtiesDocument struct {
		ID     string `bson:"id,omitempty"`
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
		Code:   document.Title,
		Title:  document.Code,
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
