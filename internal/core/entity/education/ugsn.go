package education

import "github.com/pkg/errors"

var (
	ErrUgsnTitleIsEmpty       = errors.New("ugsn: title is empty")
	ErrUgsnIDIsEmpty          = errors.New("ugsn: id is empty")
	ErrUgsnCodeIsEmpty        = errors.New("ugsn: code is empty")
	ErrUgsnSpecialityNotFound = errors.New("ugsn: specialty not found")
	ErrUgsnTitleMaxLenTitle   = errors.New("ugsn: title is more max len or empty")
	ErrUgsnLevelIDEmpty       = errors.New("ugsn: level id is empty")
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
		id string

		// code is a unique key having the format ХХ.00.00
		// where XX is any two number.
		//
		// example: 01.00.00, 11.00.00, etc
		code    string
		title   string
		levelID string
	}

	UgsnParams struct {
		ID      string
		Code    string
		Title   string
		LevelID string
	}
)

func NewUgsn(param UgsnParams) (*Ugsn, error) {
	if param.ID == "" {
		return nil, ErrUgsnIDIsEmpty
	}

	if param.Title == "" {
		return nil, ErrUgsnTitleIsEmpty
	}

	if param.Code == "" {
		return nil, ErrUgsnCodeIsEmpty
	}

	if !UgsnCodeValidate(param.Code) {
		return nil, ErrUgsnParseCode
	}

	if param.LevelID == "" {
		return nil, ErrUgsnLevelIDEmpty
	}

	return &Ugsn{
		id:      param.ID,
		title:   param.Title,
		code:    param.Code,
		levelID: param.LevelID,
	}, nil
}

func (u *Ugsn) Title() string {
	return u.title
}

func (u *Ugsn) Code() string {
	return u.code
}

func (u *Ugsn) ID() string {
	return u.id
}

func (u *Ugsn) LeveID() string {
	return u.levelID
}

func (u *Ugsn) Rename(newTitle string) error {
	if newTitle == "" {
		return ErrUgsnTitleIsEmpty
	}

	if len(newTitle) > MaxLenTitle {
		return ErrUgsnTitleMaxLenTitle
	}

	u.title = newTitle

	return nil
}

func (u *Ugsn) ChangeCode(newCode string) error {
	if newCode == "" {
		return ErrUgsnCodeIsEmpty
	}

	if !UgsnCodeValidate(newCode) {
		return ErrUgsnParseCode
	}

	u.code = newCode

	return nil
}
