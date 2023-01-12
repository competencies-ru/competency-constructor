package competencies

import (
	"regexp"

	"github.com/pkg/errors"
)

var (
	ErrParseCompetencyCode = errors.New("error parse competency code")
	ErrParseIndicatorCode  = errors.New("error parse indicator code")
)

func match(patter, code string) bool {
	ok, err := regexp.MatchString(patter, code)
	if err != nil {
		return false
	}

	return ok
}

func ValidateCodeCompetency(code string) bool {
	return match(
		`^(УК|ОПК|ПК)-([1-9]{1}[0-9]?)$`,
		code,
	)
}

func ValidateCodeIndicator(code string) bool {
	return match(
		`^(УК|ОПК|ПК)-([1-9]{1}[0-9]?\.[1-9]{1}[0-9]?)$`,
		code,
	)
}
