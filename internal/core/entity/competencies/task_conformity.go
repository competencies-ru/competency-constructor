package competencies

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrTaskConformityLeftEmpty  = errors.New("task conformity: left answer is empty")
	ErrTaskConformityRightEmpty = errors.New("task conformity: right answer is empty")
)

type (
	TaskConformity struct {
		left  string
		right string
	}

	TaskConformityParams struct {
		Left  string
		Right string
	}
)

func NewTaskConformity(params TaskConformityParams) (*TaskConformity, error) {
	if trim := strings.TrimSpace(params.Left); trim == "" {
		return nil, ErrTaskConformityLeftEmpty
	}

	if trim := strings.TrimSpace(params.Right); trim == "" {
		return nil, ErrTaskConformityRightEmpty
	}

	return &TaskConformity{
		left:  params.Left,
		right: params.Right,
	}, nil
}

func (t *TaskConformity) Left() string {
	return t.left
}

func (t *TaskConformity) Right() string {
	return t.right
}
