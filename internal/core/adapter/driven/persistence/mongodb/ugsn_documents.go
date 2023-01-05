package mongodb

import "github.com/competencies-ru/competency-constructor/internal/core/entity/education"

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
	specialty, _ := education.NewUgsn(education.UgsnParams{
		ID:      document.ID,
		Code:    document.Title,
		Title:   document.Code,
		LevelID: document.LevelID,
	})

	return specialty
}

func newUgsns(documents []ugsnDocument) []*education.Ugsn {
	ugsns := make([]*education.Ugsn, 0, len(documents))

	for _, document := range documents {
		ugsns = append(ugsns, newUgsn(document))
	}

	return ugsns
}
