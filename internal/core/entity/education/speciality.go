package education

import (
	"github.com/pkg/errors"
)

var (
	ErrSpecialtyIDIsEmpty      = errors.New("speciality: id is empty")
	ErrSpecialityTitleIsEmpty  = errors.New("speciality: title is empty")
	ErrSpecialityCodeIsEmpty   = errors.New("speciality: code is empty")
	ErrSpecialityUgsnIDIsEmpty = errors.New("speciality: ugsnID is empty")
)

func IsInvalidSpecialtyParametersError(err error) bool {
	return errors.Is(err, ErrSpecialtyIDIsEmpty) ||
		errors.Is(err, ErrSpecialityTitleIsEmpty) ||
		errors.Is(err, ErrSpecialityCodeIsEmpty) ||
		errors.Is(err, ErrSpecialityUgsnIDIsEmpty) ||
		errors.Is(err, ErrSpecialityNotMatchCode) ||
		isInvalidCodeError(err)
}

type (
	Speciality struct {
		id     string
		code   string
		title  string
		ugsnID string
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

	if param.Code == "" {
		return nil, ErrSpecialityCodeIsEmpty
	}

	if !SpecialtyCodeValidate(param.Code) {
		return nil, ErrSpecialityParseCode
	}

	if param.UgsnID == "" {
		return nil, ErrSpecialityUgsnIDIsEmpty
	}

	return &Speciality{
		id:     param.ID,
		title:  param.Title,
		code:   param.Code,
		ugsnID: param.UgsnID,
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
