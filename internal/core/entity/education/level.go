package education

import (
	"github.com/pkg/errors"
)

var (
	ErrLevelIDIsEmpty    = errors.New("level: id is empty")
	ErrLevelTitleIsEmpty = errors.New("level: title is empty")
	ErrLevelUgsnNotFound = errors.New("level: ugsn not found")
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

func (l *Level) Ugsn(code string) (*Ugsn, error) {
	value, ok := l.ugsn[code]

	if !ok {
		return nil, errors.Wrapf(
			ErrLevelUgsnNotFound,
			"get ugsn by code: %s", code)
	}

	return value, nil
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

func (l *Level) DeleteUgsn(ucode string) error {
	if _, err := l.Ugsn(ucode); err != nil {
		return err
	}

	delete(l.ugsn, ucode)

	return nil
}

func (l *Level) AllUgsn() []*Ugsn {
	result := make([]*Ugsn, 0, len(l.ugsn))

	for _, ugsn := range l.ugsn {
		tmp := ugsn
		result = append(result, tmp)
	}

	return result
}

// Specialty methods
func (l *Level) Specialty(ucode, scode string) (*Speciality, error) {
	ugsn, err := l.Ugsn(ucode)
	if err != nil {
		return nil, err
	}

	return ugsn.speciality(scode)
}

func (l *Level) AddSpecialty(ucode string, param SpecialityParams) error {
	ugsn, err := l.Ugsn(ucode)
	if err != nil {
		return err
	}

	return ugsn.addSpeciality(param)
}

func (l *Level) DeleteSpecialty(ucode, scode string) error {
	ugsn, err := l.Ugsn(ucode)
	if err != nil {
		return err
	}

	return ugsn.deleteSpecialty(scode)
}

// Program methods
func (l *Level) Program(ucode, scode, pcode string) (*Program, error) {
	specialty, err := l.Specialty(ucode, scode)
	if err != nil {
		return nil, err
	}

	return specialty.program(pcode)
}

func (l *Level) AddProgram(ucode, scode string, param ProgramParams) error {
	specialty, err := l.Specialty(ucode, scode)
	if err != nil {
		return err
	}

	return specialty.addProgram(param)
}

func (l *Level) DeleteProgram(ucode, scode, pcode string) error {
	specialty, err := l.Specialty(ucode, scode)
	if err != nil {
		return err
	}

	return specialty.deleteProgram(pcode)
}
