package specialty

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrUgsnCodeParseCode   = errors.New("ugsn: error parse code")
	ErrSpecialityParseCode = errors.New("speciality: error parse code")
	ErrCodeIsPrefixTwoZero = errors.New("code starts with two zeros")
)

type (
	ugsnCode       string
	specialityCode string
)

func newUgsnCode(code string) (ugsnCode, error) {
	if err := IsValidUgsnCode(code); err != nil {
		return "", err
	}

	return ugsnCode(code), nil
}

func newSpecialityCode(code string) (specialityCode, error) {
	if err := IsValidSpecialtyCode(code); err != nil {
		return "", err
	}

	return specialityCode(code), nil
}

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

func (s specialityCode) String() string {
	return string(s)
}

func (u ugsnCode) String() string {
	return string(u)
}
