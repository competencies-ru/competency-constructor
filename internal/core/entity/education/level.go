package education

import (
	"github.com/pkg/errors"
)

var (
	ErrLevelIDIsEmpty    = errors.New("level: id is empty")
	ErrLevelTitleIsEmpty = errors.New("level: title is empty")
)

func IsInvalidLevelParametersError(err error) bool {
	return errors.Is(err, ErrLevelIDIsEmpty) ||
		errors.Is(err, ErrLevelTitleIsEmpty)
}

type (
	// Level represents the level of the educational program.
	Level struct {
		id    string
		title string
	}

	LevelParam struct {
		ID    string
		Title string
	}
)

func NewLevel(param LevelParam) (*Level, error) {
	if param.ID == "" {
		return nil, ErrLevelIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrLevelTitleIsEmpty
	}

	return &Level{
		id:    param.ID,
		title: param.Title,
	}, nil
}

func (l *Level) ID() string {
	return l.id
}

func (l *Level) Title() string {
	return l.title
}
