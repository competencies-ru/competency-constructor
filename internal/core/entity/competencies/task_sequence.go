package competencies

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrTaskSequenceZero      = errors.New("sequence number cannot be zero")
	ErrTaskSequenceAnswer    = errors.New("answer cannot be empty")
	ErrTaskSequenceIncorrect = errors.New("sequence is incorrect")
)

type (
	TaskSequence struct {
		sequence int
		answer   string
	}

	TaskSequenceParams struct {
		Sequence int
		Answer   string
	}
)

const MaxCountSequence = 10

func NewTaskSequence(params TaskSequenceParams) (*TaskSequence, error) {
	if params.Sequence == 0 {
		return nil, ErrTaskSequenceZero
	}

	if params.Sequence == MaxCountSequence+1 {
		return nil, ErrTaskSequenceIncorrect
	}

	if trim := strings.TrimSpace(params.Answer); trim == "" {
		return nil, ErrTaskSequenceAnswer
	}

	return &TaskSequence{
		sequence: params.Sequence,
		answer:   params.Answer,
	}, nil
}

func (s *TaskSequence) Sequence() int {
	return s.sequence
}

func (s *TaskSequence) Answer() string {
	return s.answer
}