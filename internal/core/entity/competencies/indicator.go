package competencies

import "github.com/pkg/errors"

var (
	ErrIndicatorCompetencyIDIsEmpty = errors.New("indicator: competencyID is empty")
	ErrIndicatorIDIsEmpty           = errors.New("indicator: id is empty")
	ErrIndicatorTitleIsEmpty        = errors.New("indicator: title is empty")
	ErrIndicatorSubjectIDIsEmpty    = errors.New("indicator: code is invalid")
)

type (
	Indicator struct {
		id           string
		title        string
		code         string
		subjectID    string
		competencyID string
	}

	IndicatorParams struct {
		ID           string
		Title        string
		Code         string
		SubjectID    string
		CompetencyID string
	}
)

func NewIndicator(param IndicatorParams) (*Indicator, error) {
	if param.ID == "" {
		return nil, ErrIndicatorIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrIndicatorTitleIsEmpty
	}

	if !ValidateCodeIndicator(param.Code) {
		return nil, ErrParseIndicatorCode
	}

	if param.CompetencyID == "" {
		return nil, ErrIndicatorCompetencyIDIsEmpty
	}

	return &Indicator{
		id:           param.ID,
		title:        param.Title,
		code:         param.Code,
		subjectID:    param.SubjectID,
		competencyID: param.CompetencyID,
	}, nil
}

func (i *Indicator) ID() string {
	return i.id
}

func (i *Indicator) Title() string {
	return i.title
}

func (i *Indicator) Code() string {
	return i.code
}

func (i *Indicator) SubjectID() string {
	return i.subjectID
}

func (i *Indicator) CompetencyID() string {
	return i.competencyID
}

func (i *Indicator) AddSubjectID(subjectID string) error {
	if subjectID == "" {
		return ErrIndicatorSubjectIDIsEmpty
	}

	i.subjectID = subjectID

	return nil
}

func (i *Indicator) Rename(title string) error {
	if title == "" {
		return ErrIndicatorTitleIsEmpty
	}

	i.title = title

	return nil
}
