package specialty

import "github.com/pkg/errors"

var (
	ErrUgsnTitleIsEmpty       = errors.New("ugsn: title is empty")
	ErrUgsnCodeIsEmpty        = errors.New("ugsn: code is empty")
	ErrUgsnSpecialityNotFound = errors.New("ugsn: specialty not found")
)

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
		code  ugsnCode
		title string

		// Key specialityCode.
		// Value pointer Speciality
		specialities map[specialityCode]*Speciality
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

	ucode, err := newUgsnCode(param.Code)
	if err != nil {
		return nil, err
	}

	return &Ugsn{
		title:        param.Title,
		code:         ucode,
		specialities: make(map[specialityCode]*Speciality),
	}, nil
}

func (e *Ugsn) Title() string {
	return e.title
}

func (e *Ugsn) Code() string {
	return string(e.code)
}

func (e *Ugsn) AddSpeciality(s SpecialityParams) error {
	speciality, err := NewSpeciality(s)
	if err != nil {
		return err
	}

	if _, ok := e.specialities[speciality.code]; !ok {
		e.specialities[speciality.code] = speciality
	}

	return nil
}

func (e *Ugsn) Speciality(code string) (*Speciality, error) {
	scode, err := newSpecialityCode(code)
	if err != nil {
		return nil, ErrSpecialityParseCode
	}

	s, ok := e.specialities[scode]

	if !ok {
		return nil, errors.Wrap(ErrUgsnSpecialityNotFound, code)
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
