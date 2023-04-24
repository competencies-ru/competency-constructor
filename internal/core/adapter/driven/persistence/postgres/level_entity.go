package postgres

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type levelEntity struct {
	ID    string
	Title string
}

func newLevel(entity levelEntity) *education.Level {
	level, _ := education.NewLevel(education.LevelParam{
		ID:    entity.ID,
		Title: entity.Title,
	})

	return level
}

func newLevelModel(document levelEntity) query.LevelModel {
	return query.LevelModel{
		ID:    document.ID,
		Title: document.Title,
	}
}

func newLevelModels(documents []levelEntity) []query.LevelModel {
	levels := make([]query.LevelModel, 0, len(documents))

	for _, document := range documents {
		levels = append(levels, newLevelModel(document))
	}

	return levels
}
