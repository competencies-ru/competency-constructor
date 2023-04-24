package postgres

import (
	"database/sql"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

type competencyEntity struct {
	ID             string
	Code           string
	Title          string
	Category       string
	CompetencyType competencies.Type
	LevelID        sql.NullString
	UgsnID         sql.NullString
	SpecialtyID    sql.NullString
	ProgramID      sql.NullString
}

func newCompetency(entity competencyEntity) *competencies.Competency {
	cmp, _ := competencies.NewCompetency(competencies.CompetencyParam{
		ID:             entity.ID,
		Title:          entity.Title,
		Code:           entity.Code,
		Category:       entity.Category,
		CompetencyType: entity.CompetencyType,
		LevelID:        entity.LevelID.String,
		UgsnID:         entity.UgsnID.String,
		SpecialtyID:    entity.SpecialtyID.String,
		ProgramID:      entity.ProgramID.String,
	})

	return cmp
}

func newCompetencyEntity(cmp *competencies.Competency) competencyEntity {
	return competencyEntity{
		ID:             cmp.ID(),
		Code:           cmp.Code(),
		Title:          cmp.Title(),
		Category:       cmp.Category(),
		CompetencyType: cmp.CompetencyType(),
		LevelID:        stringToNullString(cmp.LevelID()),
		UgsnID:         stringToNullString(cmp.UgsnID()),
		SpecialtyID:    stringToNullString(cmp.SpecialtyID()),
		ProgramID:      stringToNullString(cmp.ProgramID()),
	}
}

func newCompetencies(entities []competencyEntity) []*competencies.Competency {
	result := make([]*competencies.Competency, 0, len(entities))

	for _, document := range entities {
		result = append(result, newCompetency(document))
	}

	return result
}

func newCompetencyModels(entities []competencyEntity) []query.CompetencyModel {
	result := make([]query.CompetencyModel, 0, len(entities))

	for _, document := range entities {
		result = append(result, newCompetencyModel(document))
	}

	return result
}

func newCompetencyModel(entity competencyEntity) query.CompetencyModel {
	return query.CompetencyModel{
		ID:             entity.ID,
		Code:           entity.Code,
		Title:          entity.Title,
		Category:       entity.Title,
		CompetencyType: entity.CompetencyType,
		LevelID:        entity.LevelID.String,
		UgsnID:         entity.UgsnID.String,
		SpecialtyID:    entity.SpecialtyID.String,
		ProgramID:      entity.ProgramID.String,
	}
}
