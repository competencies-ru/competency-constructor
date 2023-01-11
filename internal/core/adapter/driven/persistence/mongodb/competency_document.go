package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

type competencyDocument struct {
	ID             string            `bson:"_id"`
	Code           string            `bson:"code"`
	Title          string            `bson:"title"`
	Category       string            `bson:"category"`
	CompetencyType competencies.Type `bson:"type"`
	LevelID        string            `bson:"level_id"`
	UgsnID         string            `bson:"ugsn_id"`
	SpecialtyID    string            `bson:"specialty_id"`
	ProgramID      string            `bson:"program_id"`
}

func newCompetencyDocument(cmp *competencies.Competency) competencyDocument {
	return competencyDocument{
		ID:             cmp.ID(),
		Code:           cmp.Code(),
		Title:          cmp.Title(),
		Category:       cmp.Category(),
		CompetencyType: cmp.CompetencyType(),
		LevelID:        cmp.LevelID(),
		UgsnID:         cmp.UgsnID(),
		SpecialtyID:    cmp.SpecialtyID(),
		ProgramID:      cmp.ProgramID(),
	}
}

func newCompetency(document competencyDocument) *competencies.Competency {
	cmp, _ := competencies.NewCompetency(competencies.CompetencyParam{
		ID:             document.ID,
		Title:          document.Title,
		Code:           document.Code,
		Category:       document.Category,
		CompetencyType: document.CompetencyType,
		LevelID:        document.LevelID,
		UgsnID:         document.UgsnID,
		SpecialtyID:    document.SpecialtyID,
		ProgramID:      document.ProgramID,
	})

	return cmp
}

func newCompetencies(documents []competencyDocument) []*competencies.Competency {
	result := make([]*competencies.Competency, 0, len(documents))

	for _, document := range documents {
		result = append(result, newCompetency(document))
	}

	return result
}

func newCompetencyModels(documents []competencyDocument) []query.CompetencyModel {
	result := make([]query.CompetencyModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, newCompetencyModel(document))
	}

	return result
}

func newCompetencyModel(document competencyDocument) query.CompetencyModel {
	return query.CompetencyModel{
		ID:             document.ID,
		Code:           document.Code,
		Title:          document.Title,
		Category:       document.Title,
		CompetencyType: document.CompetencyType,
		LevelID:        document.LevelID,
		UgsnID:         document.UgsnID,
		SpecialtyID:    document.SpecialtyID,
		ProgramID:      document.ProgramID,
	}
}
