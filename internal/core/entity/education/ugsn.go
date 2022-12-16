package education

import "github.com/pkg/errors"

var (
	ErrUgsnTitleIsEmpty       = errors.New("ugsn: title is empty")
	ErrUgsnCodeIsEmpty        = errors.New("ugsn: code is empty")
	ErrUgsnSpecialityNotFound = errors.New("ugsn: education not found")
	ErrUgsnTitleMaxLenTitle   = errors.New("ugsn: title is more max len or empty")
)

func IsInvalidUgsnParametersError(err error) bool {
	return errors.Is(err, ErrUgsnTitleIsEmpty) ||
		errors.Is(err, ErrUgsnCodeIsEmpty) ||
		errors.Is(err, ErrUgsnSpecialityNotFound) ||
		errors.Is(err, ErrUgsnTitleMaxLenTitle) ||
		isInvalidCodeError(err)
}

const MaxLenTitle = 1000

type (

	// Ugsn is an enlarged
	// group of specialties and directions
	//
	// Example: 09.00.00
	// Информатика и вычислительная техника.
	Ugsn struct {
		// code is a unique key having the format ХХ.00.00
		// where XX is any two number.
		//
		// example: 01.00.00, 11.00.00, etc
		code  string
		title string

		// Key specialityCode.
		// Value pointer Speciality
		specialities map[string]*Speciality
	}

	UgsnParams struct {
		Code  string
		Title string
	}
)

func NewUgsn(param UgsnParams) (*Ugsn, error) {
	if param.Title == "" {
		return nil, ErrUgsnTitleIsEmpty
	}

	if param.Code == "" {
		return nil, ErrUgsnCodeIsEmpty
	}

	if err := IsValidUgsnCode(param.Code); err != nil {
		return nil, err
	}

	return &Ugsn{
		title:        param.Title,
		code:         param.Code,
		specialities: make(map[string]*Speciality),
	}, nil
}

func (e *Ugsn) Title() string {
	return e.title
}

func (e *Ugsn) Code() string {
	return e.code
}

func (e *Ugsn) AddSpeciality(s SpecialityParams) error {
	speciality, err := NewSpeciality(s)
	if err != nil {
		return errors.Wrapf(err, "adding education by code: %s", s.Code)
	}

	if _, ok := e.specialities[speciality.code]; !ok {
		e.specialities[speciality.code] = speciality
	}

	return nil
}

func (e *Ugsn) Speciality(code string) (*Speciality, error) {
	s, ok := e.specialities[code]
	if !ok {
		return nil, errors.Wrapf(
			ErrUgsnSpecialityNotFound,
			"get education by code: %s", code)
	}

	return s, nil
}

func (e *Ugsn) Specialities() []*Speciality {
	specialities := make([]*Speciality, 0, len(e.specialities))
	for _, value := range e.specialities {
		specialities = append(specialities, value)
	}

	return specialities
}

func (e *Ugsn) Rename(newTitle string) error {
	if newTitle == "" || len(newTitle) > MaxLenTitle {
		return ErrUgsnTitleMaxLenTitle
	}

	e.title = newTitle

	return nil
}

func (e *Ugsn) DeleteSpecialty(code string) error {
	if _, ok := e.specialities[code]; !ok {
		return errors.Wrapf(
			ErrUgsnSpecialityNotFound,
			"get education by code: %s", code)
	}

	delete(e.specialities, code)

	return nil
}
