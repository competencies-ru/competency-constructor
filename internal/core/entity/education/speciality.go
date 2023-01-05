package education

import (
	"github.com/pkg/errors"
)

var (
	ErrSpecialtyIDIsEmpty        = errors.New("speciality: id is empty")
	ErrSpecialityTitleIsEmpty    = errors.New("speciality: title is empty")
	ErrSpecialityCodeIsEmpty     = errors.New("speciality: code is empty")
	ErrSpecialityUgsnCodeIsEmpty = errors.New("speciality: ugsnCode is empty")
	ErrSpecialityProgramNotFound = errors.New("speciality: program not found")
)

func IsInvalidSpecialtyParametersError(err error) bool {
	return errors.Is(err, ErrSpecialtyIDIsEmpty) ||
		errors.Is(err, ErrSpecialityTitleIsEmpty) ||
		errors.Is(err, ErrSpecialityCodeIsEmpty) ||
		errors.Is(err, ErrSpecialityUgsnCodeIsEmpty) ||
		errors.Is(err, ErrSpecialityNotMatchCode) ||
		isInvalidCodeError(err)
}

type (
	Speciality struct {
		id     string
		code   string
		title  string
		ugsnID string
		// key programCode, value pointer Program
		programs map[string]*Program
	}

	SpecialityParams struct {
		ID     string
		Code   string
		Title  string
		UgsnID string
	}
)

func NewSpeciality(param SpecialityParams) (*Speciality, error) {
	if param.ID == "" {
		return nil, ErrSpecialtyIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrSpecialityTitleIsEmpty
	}

	if SpecialtyCodeValidate(param.Code) {
		return nil, ErrSpecialityParseCode
	}

	return &Speciality{
		id:       param.ID,
		title:    param.Title,
		code:     param.Code,
		ugsnID:   param.UgsnID,
		programs: make(map[string]*Program),
	}, nil
}

func (s *Speciality) ID() string {
	return s.id
}

func (s *Speciality) Title() string {
	return s.title
}

func (s *Speciality) Code() string {
	return s.code
}

func (s *Speciality) UgsnID() string {
	return s.ugsnID
}

func (s *Speciality) AddProgram(p ProgramParams) error {
	program, err := NewProgram(p)
	if err != nil {
		return err
	}

	if _, ok := s.programs[program.code]; !ok {
		s.programs[program.code] = program
	}

	return nil
}

func (s *Speciality) ProgramByCode(code string) (*Program, error) {
	p, ok := s.programs[code]

	if !ok {
		return nil, errors.Wrapf(ErrSpecialityProgramNotFound, "get program by code %s", code)
	}

	return p, nil
}

func (s *Speciality) ProgramByID(id string) (*Program, error) {
	for _, program := range s.programs {
		if program.ID() == id {
			return program, nil
		}
	}

	return nil, ErrSpecialityProgramNotFound
}

func (s *Speciality) Programs() []*Program {
	programs := make([]*Program, 0, len(s.programs))
	for _, value := range s.programs {
		programs = append(programs, value)
	}

	return programs
}
