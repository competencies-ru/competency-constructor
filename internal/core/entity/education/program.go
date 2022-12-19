package education

import "github.com/pkg/errors"

const maxLenTitleProgram = 1000

var (
	ErrProgramIDIsEmpty            = errors.New("program: id is empty")
	ErrProgramTitleIsEmpty         = errors.New("program: title is empty")
	ErrProgramSpecialtyCodeIsEmpty = errors.New("program: specialityCode is empty")
	ErrProgramMaxLenTitle          = errors.New("program: title is is more max len")
)

type (
	Program struct {
		code           string
		title          string
		specialityCode string
	}

	ProgramParams struct {
		Code          string
		Title         string
		SpecialtyCode string
	}
)

func NewProgram(param ProgramParams) (*Program, error) {
	if param.Code == "" {
		return nil, ErrProgramIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrProgramTitleIsEmpty
	}

	if param.SpecialtyCode == "" {
		return nil, ErrProgramSpecialtyCodeIsEmpty
	}

	if err := IsValidSpecialtyCode(param.SpecialtyCode); err != nil {
		return nil, err
	}

	if err := IsValidProgramCode(param.Code); err != nil {
		return nil, err
	}

	return &Program{code: param.Code, title: param.Title, specialityCode: param.SpecialtyCode}, nil
}

func (p *Program) Title() string {
	return p.title
}

func (p *Program) Code() string {
	return p.code
}

func (p *Program) SpecialityCode() string {
	return p.specialityCode
}

func (p *Program) Rename(title string) error {
	if len(title) > maxLenTitleProgram {
		return ErrProgramMaxLenTitle
	}

	p.title = title

	return nil
}
