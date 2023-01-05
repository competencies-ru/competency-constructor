package education

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrUgsnParseCode          = errors.New("ugsn: error parse code")
	ErrSpecialityParseCode    = errors.New("speciality: error parse code")
	ErrProgramParseCode       = errors.New("program: error parse code")
	ErrSpecialityNotMatchCode = errors.New("speciality: code does not match ugsnCode")
	ErrProgramNotMatchCode    = errors.New("program: code does not match specialty code")
)

func match(patter, code string) bool {
	ok, err := regexp.MatchString(patter, code)
	if err != nil {
		return false
	}

	return ok
}

func UgsnCodeValidate(code string) bool {
	return match(`^(0[1-9]{1}|[1-9]{1}[0-9]{1})\.00\.00$`, code)
}

func SpecialtyCodeValidate(code string) bool {
	return match(
		`^(0[1-9]{1}|[1-9]{1}[0-9]{1})\.(0[1-9]{1}|[1-9]{1}[0-9]{1})\.(0[1-9]{1}|[1-9]{1}[0-9]{1})$`,
		code,
	)
}

func ProgramCodeValidate(code string) bool {
	return match(
		`^(0[1-9]{1}|[1-9]{1}[0-9]{1})\.(0[1-9]{1}|[1-9]{1}[0-9]{1})\.(0[1-9]{1}|[1-9]{1}[0-9]{1})-(0[1-9]{1}|[1-9]{1}[0-9]{1})$`,
		code,
	)
}

func isInvalidCodeError(err error) bool {
	return errors.Is(err, ErrProgramParseCode) ||
		errors.Is(err, ErrUgsnParseCode) ||
		errors.Is(err, ErrSpecialityParseCode)
}

func MatchProgramCode(pcode, scode string) bool {
	return strings.Contains(pcode[:8], scode[:8])
}

func MatchSpecialtyCode(scode, ucode string) bool {
	return strings.Contains(scode[:2], ucode[:2])
}
