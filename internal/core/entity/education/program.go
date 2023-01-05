package education

import (
	"github.com/pkg/errors"
)

const maxLenTitleProgram = 1000

var (
	ErrProgramIDIsEmpty            = errors.New("program: id is empty")
	ErrProgramCodeIsEmpty          = errors.New("program: code is empty")
	ErrProgramTitleIsEmpty         = errors.New("program: title is empty")
	ErrProgramSpecialtyCodeIsEmpty = errors.New("program: specialityCode is empty")
	ErrProgramMaxLenTitle          = errors.New("program: title is is more max len")
)

type (
	Program struct {
		id          string
		code        string
		title       string
		specialtyID string
	}

	ProgramParams struct {
		ID            string
		Code          string
		Title         string
		SpecialtyCode string
		SpecialtyID   string
	}
)

func NewProgram(param ProgramParams) (*Program, error) {
	if param.ID == "" {
		return nil, ErrProgramIDIsEmpty
	}

	if param.Code == "" {
		return nil, ErrProgramCodeIsEmpty
	}

	if param.Title == "" {
		return nil, ErrProgramTitleIsEmpty
	}

	if SpecialtyCodeValidate(param.SpecialtyCode) {
		return nil, ErrSpecialityParseCode
	}

	if ProgramCodeValidate(param.Code) {
		return nil, ErrProgramParseCode
	}

	if !matchProgramCode(param.Code, param.SpecialtyCode) {
		return nil, ErrProgramNotMatchCode
	}

	return &Program{id: param.ID, code: param.Code, title: param.Title, specialtyID: param.SpecialtyID}, nil
}

func (p *Program) ID() string {
	return p.id
}

func (p *Program) Title() string {
	return p.title
}

func (p *Program) Code() string {
	return p.code
}

func (p *Program) SpecialtyID() string {
	return p.specialtyID
}

func (p *Program) Rename(title string) error {
	if len(title) > maxLenTitleProgram {
		return ErrProgramMaxLenTitle
	}

	p.title = title

	return nil
}
