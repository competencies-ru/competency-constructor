package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	levelDocument struct {
		ID    string `bson:"_id,omitempty"`
		Title string `bson:"title,omitempty"`
	}
)

func newLevel(document levelDocument) *education.Level {
	level, _ := education.NewLevel(education.LevelParam{
		ID:    document.ID,
		Title: document.Title,
	})

	return level
}

func newLevels(documents []levelDocument) []*education.Level {
	levels := make([]*education.Level, 0, len(documents))

	for _, document := range documents {
		levels = append(levels, newLevel(document))
	}

	return levels
}

func newLevelDocument(level *education.Level) levelDocument {
	return levelDocument{
		ID:    level.ID(),
		Title: level.Title(),
	}
}
