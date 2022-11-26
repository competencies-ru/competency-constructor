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
	if ok := match(`^\d{2}.0{2}.0{2}$`, code); !ok {
		return "", ErrUgsnCodeParseCode
	}

	if err := isPrefixTwoZero(code); err != nil {
		return "", err
	}

	return ugsnCode(code), nil
}

func newSpecialityCode(code string) (specialityCode, error) {
	if ok := match(`^\d{2}.\d{2}.\d{2}$`, code); !ok {
		return "", ErrSpecialityParseCode
	}

	if err := isPrefixTwoZero(code); err != nil {
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

func isPrefixTwoZero(code string) error {
	if strings.HasPrefix(code, "00") {
		return ErrCodeIsPrefixTwoZero
	}

	return nil
}
