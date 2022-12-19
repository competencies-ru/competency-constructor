package education

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrUgsnCodeParseCode   = errors.New("ugsn: error parse code")
	ErrSpecialityParseCode = errors.New("speciality: error parse code")
	ErrProgramParseCode    = errors.New("program: error parse code")
	ErrCodeIsPrefixTwoZero = errors.New("code starts with two zeros")
)

func match(patter, code string) bool {
	ok, err := regexp.MatchString(patter, code)
	if err != nil {
		return false
	}

	return ok
}

func isPrefixTwoZero(code string) bool {
	return strings.HasPrefix(code, "00")
}

func IsValidUgsnCode(code string) error {
	if isPrefixTwoZero(code) {
		return ErrCodeIsPrefixTwoZero
	}

	if !match(`^\d{2}.0{2}.0{2}$`, code) {
		return ErrUgsnCodeParseCode
	}

	return nil
}

func IsValidSpecialtyCode(code string) error {
	if isPrefixTwoZero(code) {
		return ErrCodeIsPrefixTwoZero
	}

	if match(`^\d{2}.0{2}.0{2}$`, code) || !match(`^\d{2}.\d{2}.\d{2}$`, code) {
		return ErrSpecialityParseCode
	}

	return nil
}

func IsValidProgramCode(code string) error {
	if isPrefixTwoZero(code) {
		return ErrCodeIsPrefixTwoZero
	}

	if match(`^\d{2}.0{2}.0{2}$`, code) || !match(`^\d{2}\.\d{2}\.\d{2}-\d{2}$`, code) {
		return ErrProgramParseCode
	}

	return nil
}

func isInvalidCodeError(err error) bool {
	return errors.Is(err, ErrCodeIsPrefixTwoZero) ||
		errors.Is(err, ErrUgsnCodeParseCode) ||
		errors.Is(err, ErrSpecialityParseCode)
}
