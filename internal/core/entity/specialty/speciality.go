package specialty

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
		code     specialityCode
		title    string
		ugsnCode ugsnCode

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

	scode, err := newSpecialityCode(param.Code)
	if err != nil {
		return nil, err
	}

	if param.UgsnCode == "" {
		return nil, ErrSpecialityUgsnCodeIsEmpty
	}

	ucode, err := newUgsnCode(param.UgsnCode)
	if err != nil {
		return nil, err
	}

	if !validateCodes(scode, ucode) {
		return nil, ErrSpecialityNotMatchCode
	}

	return &Speciality{
		title:    param.Title,
		code:     scode,
		ugsnCode: ucode,
		programs: make(map[string]*Program),
	}, nil
}

func (s *Speciality) Title() string {
	return s.title
}

func (s *Speciality) Code() string {
	return string(s.code)
}

func (s *Speciality) UgsnCode() string {
	return string(s.ugsnCode)
}

func (s *Speciality) AddProgram(p ProgramParams) error {
	program, err := NewProgram(p)
	if err != nil {
		return err
	}

	if _, ok := s.programs[program.id]; !ok {
		s.programs[program.id] = program
	}

	return nil
}

func validateCodes(scode specialityCode, ucode ugsnCode) bool {
	return strings.Contains(string(scode)[:2], string(ucode)[:2])
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
