package postgres

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type ugsnEntity struct {
	ID      string
	Code    string
	Title   string
	LevelID string
}

func newUgsnEntity(ugsn *education.Ugsn) ugsnEntity {
	return ugsnEntity{
		ID:      ugsn.ID(),
		Code:    ugsn.Code(),
		Title:   ugsn.Title(),
		LevelID: ugsn.LeveID(),
	}
}

func newUgsn(entity ugsnEntity) *education.Ugsn {
	ugsn, _ := education.NewUgsn(education.UgsnParams{
		ID:      entity.ID,
		Code:    entity.Code,
		Title:   entity.Title,
		LevelID: entity.LevelID,
	})

	return ugsn
}

func newUgsnModels(documents []ugsnEntity) []query.UgsnModel {
	result := make([]query.UgsnModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.UgsnModel{
			ID:      document.ID,
			Code:    document.Code,
			Title:   document.Title,
			LevelID: document.LevelID,
		})
	}

	return result
}
