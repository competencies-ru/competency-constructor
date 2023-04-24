package postgres

import (
	"database/sql"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

type indicatorEntity struct {
	ID           string
	Title        string
	Code         string
	CompetencyID string
	SubjectID    sql.NullString

	Subject subjectEntity
}

func newIndicatorModels(entities []indicatorEntity) []query.IndicatorModel {
	buff := make([]query.IndicatorModel, 0, len(entities))

	for _, entity := range entities {
		buff = append(buff, newIndicatorModel(entity))
	}

	return buff
}

func newIndicatorModel(entity indicatorEntity) query.IndicatorModel {
	indicator := query.IndicatorModel{
		ID:           entity.ID,
		Code:         entity.Code,
		Title:        entity.Title,
		CompetencyID: entity.CompetencyID,
	}

	if !subjectIsEmpty(entity.Subject) {
		subjectModel := newSubjectModel(entity.Subject)
		indicator.Subject = &subjectModel
	}

	return indicator
}

func newIndicator(entity indicatorEntity) *competencies.Indicator {
	indicator, _ := competencies.NewIndicator(competencies.IndicatorParams{
		ID:           entity.ID,
		Title:        entity.Title,
		Code:         entity.Code,
		SubjectID:    entity.SubjectID.String,
		CompetencyID: entity.CompetencyID,
	})

	return indicator
}
