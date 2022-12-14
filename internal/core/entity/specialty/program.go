package specialty

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
		id             string
		title          string
		specialityCode specialityCode
	}

	ProgramParams struct {
		ID            string
		Title         string
		SpecialtyCode string
	}
)

func NewProgram(param ProgramParams) (*Program, error) {
	if param.ID == "" {
		return nil, ErrProgramIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrProgramTitleIsEmpty
	}

	if param.SpecialtyCode == "" {
		return nil, ErrProgramSpecialtyCodeIsEmpty
	}

	scode, err := newSpecialityCode(param.SpecialtyCode)
	if err != nil {
		return nil, err
	}

	return &Program{id: param.ID, title: param.Title, specialityCode: scode}, nil
}

func (p *Program) Title() string {
	return p.title
}

func (p *Program) ID() string {
	return p.id
}

func (p *Program) SpecialityCode() string {
	return string(p.specialityCode)
}

func (p *Program) Rename(title string) error {
	if len(title) > maxLenTitleProgram {
		return ErrProgramMaxLenTitle
	}

	p.title = title

	return nil
}
