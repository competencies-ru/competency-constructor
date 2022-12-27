package education

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrSpecialtyIDIsEmpty        = errors.New("speciality: id is empty")
	ErrSpecialityTitleIsEmpty    = errors.New("speciality: title is empty")
	ErrSpecialityCodeIsEmpty     = errors.New("speciality: code is empty")
	ErrSpecialityUgsnCodeIsEmpty = errors.New("speciality: ugsnCode is empty")
	ErrSpecialityNotMatchCode    = errors.New("speciality: code does not match ugsnCode")
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
		id    string
		code  string
		title string
		// key programCode, value pointer Program
		programs map[string]*Program
	}

	SpecialityParams struct {
		ID       string
		Code     string
		Title    string
		UgsnCode string
	}
)

func NewSpeciality(param SpecialityParams) (*Speciality, error) {
	if param.ID == "" {
		return nil, ErrSpecialtyIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrSpecialityTitleIsEmpty
	}

	if param.Code == "" {
		return nil, ErrSpecialityCodeIsEmpty
	}

	if err := IsValidSpecialtyCode(param.Code); err != nil {
		return nil, err
	}

	if param.UgsnCode == "" {
		return nil, ErrSpecialityUgsnCodeIsEmpty
	}

	if err := IsValidUgsnCode(param.UgsnCode); err != nil {
		return nil, err
	}

	if !validateCodes(param.Code, param.UgsnCode) {
		return nil, ErrSpecialityNotMatchCode
	}

	return &Speciality{
		id:       param.ID,
		title:    param.Title,
		code:     param.Code,
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

func (s *Speciality) addProgram(p ProgramParams) error {
	program, err := NewProgram(p)
	if err != nil {
		return err
	}

	if _, ok := s.programs[program.code]; !ok {
		s.programs[program.code] = program
	}

	return nil
}

func (s *Speciality) deleteProgram(pcode string) error {
	if _, ok := s.programs[pcode]; !ok {
		return errors.Wrapf(
			ErrUgsnSpecialityNotFound,
			"get program by code: %s", pcode)
	}

	delete(s.programs, pcode)

	return nil
}

func validateCodes(scode string, ucode string) bool {
	return strings.Contains(scode[:2], ucode[:2])
}

func (s *Speciality) program(code string) (*Program, error) {
	p, ok := s.programs[code]

	if !ok {
		return nil, errors.Wrapf(ErrSpecialityProgramNotFound, "get program by code %s", code)
	}

	return p, nil
}

func (s *Speciality) Programs() []*Program {
	programs := make([]*Program, 0, len(s.programs))
	for _, value := range s.programs {
		programs = append(programs, value)
	}

	return programs
}
