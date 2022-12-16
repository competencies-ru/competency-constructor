package education

import "github.com/pkg/errors"

var (
	ErrLevelIDIsEmpty    = errors.New("level: id is empty")
	ErrLevelTitleIsEmpty = errors.New("level: title is empty")
)

type (
	// Level represents the level of the educational program.
	Level struct {
		id    string
		title string

		ugsn map[string]*Ugsn
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
		ugsn:  make(map[string]*Ugsn),
	}, nil
}

func (l *Level) ID() string {
	return l.id
}

func (l *Level) Title() string {
	return l.title
}

func (l *Level) AddUgsn(param UgsnParams) error {
	if _, ok := l.ugsn[param.Code]; ok {
		return nil
	}

	ugsn, err := NewUgsn(param)
	if err != nil {
		return err
	}

	l.ugsn[ugsn.Code()] = ugsn

	return nil
}
