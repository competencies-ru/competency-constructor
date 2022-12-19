package education

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrSpecialityTitleIsEmpty    = errors.New("speciality: title is empty")
	ErrSpecialityCodeIsEmpty     = errors.New("speciality: code is empty")
	ErrSpecialityUgsnCodeIsEmpty = errors.New("speciality: ugsnCode is empty")
	ErrSpecialityNotMatchCode    = errors.New("speciality: code does not match ugsnCode")
	ErrSpecialityProgramNotFound = errors.New("speciality: program not found")
)

type (
	Speciality struct {
		code     string
		title    string
		ugsnCode string

		// key programID, value pointer Program
		programs map[string]*Program
	}

	SpecialityParams struct {
		Code     string
		Title    string
		UgsnCode string
	}
)

func NewSpeciality(param SpecialityParams) (*Speciality, error) {
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
		title:    param.Title,
		code:     param.Code,
		ugsnCode: param.UgsnCode,
		programs: make(map[string]*Program),
	}, nil
}

func (s *Speciality) Title() string {
	return s.title
}

func (s *Speciality) Code() string {
	return s.code
}

func (s *Speciality) UgsnCode() string {
	return s.ugsnCode
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

func validateCodes(scode string, ucode string) bool {
	return strings.Contains(scode[:2], ucode[:2])
}

func (s *Speciality) Program(id string) (*Program, error) {
	p, ok := s.programs[id]

	if !ok {
		return nil, errors.Wrap(ErrSpecialityProgramNotFound, id)
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
