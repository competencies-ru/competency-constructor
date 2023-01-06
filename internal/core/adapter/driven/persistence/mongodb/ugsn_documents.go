package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	ugsnDocument struct {
		ID      string `bson:"_id,omitempty"`
		Code    string `bson:"code"`
		Title   string `bson:"title"`
		LevelID string `bson:"level_id"`
	}
)

func newUgsnDocument(ugsn *education.Ugsn) ugsnDocument {
	return ugsnDocument{
		ID:      ugsn.ID(),
		Code:    ugsn.Code(),
		Title:   ugsn.Title(),
		LevelID: ugsn.LeveID(),
	}
}

func newUgsn(document ugsnDocument) *education.Ugsn {
	ugsn, _ := education.NewUgsn(education.UgsnParams{
		ID:      document.ID,
		Code:    document.Code,
		Title:   document.Title,
		LevelID: document.LevelID,
	})

	return ugsn
}

func newUgsns(documents []ugsnDocument) []*education.Ugsn {
	ugsns := make([]*education.Ugsn, 0, len(documents))

	for _, document := range documents {
		ugsns = append(ugsns, newUgsn(document))
	}

	return ugsns
}

func newUgsnModels(documents []ugsnDocument) []query.UgsnModel {
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
