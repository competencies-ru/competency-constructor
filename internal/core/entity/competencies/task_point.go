package competencies

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrTaskPointVariantsIsEmpty       = errors.New("task point: number of answer options is less than one")
	ErrTaskPointVariantsIsEmptyValue  = errors.New("task point: one or more variants options is blank")
	ErrTaskPointVariantsMoreMax       = errors.New("task point: variants more maximum")
	ErrTaskPointAnswerIncorrect       = errors.New("task point: one or more answer options is incorrect")
	ErrTaskPointAnswerIsEmpty         = errors.New("task point: answer is empty")
	ErrTaskPointAnswerNumberIncorrect = errors.New("task point: incorrect number answer")
)

const (
	MaxAnswerNumber = 3
	MaxVariantsLen  = 6
)

type (
	TaskPoint struct {
		single   bool
		variants []string
		answers  []int
	}

	TaskPointParams struct {
		Single   bool
		Variants []string
		Answers  []int
	}
)

func NewTaskPoint(params TaskPointParams) (*TaskPoint, error) {
	if err := validateVariants(params.Variants); err != nil {
		return nil, err
	}

	if err := validateAnswers(params); err != nil {
		return nil, err
	}

	return &TaskPoint{
		single:   params.Single,
		variants: params.Variants,
		answers:  params.Answers,
	}, nil
}

func (t *TaskPoint) IsSingle() bool {
	return t.single
}

func (t *TaskPoint) Variants() []string {
	return t.variants
}

func (t *TaskPoint) Answer() []int {
	return t.answers
}

func validateVariants(variants []string) error {
	if len(variants) <= 1 {
		return ErrTaskPointVariantsIsEmpty
	}

	if len(variants) > MaxVariantsLen {
		return ErrTaskPointVariantsMoreMax
	}

	for _, variant := range variants {
		if trim := strings.TrimSpace(variant); trim == "" {
			return ErrTaskPointVariantsIsEmptyValue
		}
	}

	return nil
}

func validateAnswers(params TaskPointParams) error {
	if len(params.Answers) < 1 {
		return ErrTaskPointAnswerIsEmpty
	}

	maxIndexVariants := len(params.Variants) - 1

	if params.Single {

		if len(params.Answers) != 1 {
			return ErrTaskPointAnswerNumberIncorrect
		}

		if params.Answers[0] < 0 || params.Answers[0] > maxIndexVariants {
			return ErrTaskPointAnswerIncorrect
		}

		return nil
	}

	if len(params.Answers) < 2 || len(params.Answers) > MaxAnswerNumber {
		return ErrTaskPointAnswerNumberIncorrect
	}

	for _, answer := range params.Answers {
		if answer < 0 || answer > maxIndexVariants {
			return ErrTaskPointAnswerIncorrect
		}
	}

	return nil
}
