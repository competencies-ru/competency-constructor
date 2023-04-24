package postgres

import "github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"

type (
	questionEntity struct {
		ID             string
		Description    string
		IndicatorID    string
		Rank           int
		QuestionType   competencies.QuestionType
		CompleteAnswer string
	}

	taskConformity struct {
		Left  string
		Right string
	}

	taskPointEntity struct {
		Variants []string
		Answers  []int
		Single   bool
	}

	taskSequenceEntity struct {
		Answer   string
		Sequence int
	}
)
