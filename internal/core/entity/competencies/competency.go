package competencies

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrProgramIDIsEmpty            = errors.New("competency: programID is empty")
	ErrLevelIDIsEmpty              = errors.New("competency: levelID is empty")
	ErrSpecialtyIDAndUgsnIDIsEmpty = errors.New("competency: ugsnID and specialtyID is empty")
	ErrCompetencyIDIsEmpty         = errors.New("competency: id is empty")
	ErrCompetencyTitleIsEmpty      = errors.New("competency: title is empty")
	ErrCompetencyCategoryIsEmpty   = errors.New("competency: category is empty")
	ErrCompetencyTypeInvalid       = errors.New("competency: type is invalid")
	ErrCompetencyCodeInvalid       = errors.New("competency: code is invalid")
)

type (
	Competency struct {
		id             string
		title          string
		code           string
		category       string
		levelID        string
		ugsnID         string
		specialtyID    string
		programID      string
		competencyType Type
	}

	CompetencyParam struct {
		ID             string
		Title          string
		Code           string
		Category       string
		LevelID        string
		UgsnID         string
		SpecialtyID    string
		ProgramID      string
		CompetencyType Type
	}
)

func NewCompetency(param CompetencyParam) (*Competency, error) {
	if param.ID == "" {
		return nil, ErrCompetencyIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrCompetencyTitleIsEmpty
	}

	if !ValidateCodeCompetency(param.Code) {
		return nil, ErrParseCompetencyCode
	}

	if param.Category == "" {
		return nil, ErrCompetencyCategoryIsEmpty
	}

	if !param.CompetencyType.IsValid() {
		return nil, ErrCompetencyTypeInvalid
	}

	if !isMatchCode(param.Code, param.CompetencyType) {
		return nil, ErrCompetencyCodeInvalid
	}

	if err := validateEducationParameters(param); err != nil {
		return nil, err
	}

	return &Competency{
		id:             param.ID,
		title:          param.Title,
		code:           param.Code,
		category:       param.Category,
		competencyType: param.CompetencyType,
		levelID:        param.LevelID,
		ugsnID:         param.UgsnID,
		specialtyID:    param.SpecialtyID,
		programID:      param.ProgramID,
	}, nil
}

func isMatchCode(code string, p Type) bool {
	switch p {
	case GENERAL:
		return strings.Contains(code, p.String())
	case PROFESSIONAL:
		return strings.Contains(code, p.String())
	case UNIVERSAL:
		return strings.Contains(code, p.String())
	}

	return false
}

func validateEducationParameters(param CompetencyParam) error {
	if param.CompetencyType == UNIVERSAL && param.LevelID == "" {
		return ErrLevelIDIsEmpty
	}

	if param.CompetencyType == PROFESSIONAL && param.ProgramID == "" {
		return ErrProgramIDIsEmpty
	}

	if param.CompetencyType == GENERAL && param.UgsnID == "" && param.SpecialtyID == "" {
		return ErrSpecialtyIDAndUgsnIDIsEmpty
	}

	return nil
}

func (c Competency) ID() string {
	return c.id
}

func (c Competency) Title() string {
	return c.title
}

func (c Competency) Code() string {
	return c.code
}

func (c Competency) Category() string {
	return c.category
}

func (c Competency) CompetencyType() Type {
	return c.competencyType
}

func (c Competency) LevelID() string {
	return c.levelID
}

func (c Competency) UgsnID() string {
	return c.ugsnID
}

func (c Competency) SpecialtyID() string {
	return c.specialtyID
}

func (c Competency) ProgramID() string {
	return c.programID
}
