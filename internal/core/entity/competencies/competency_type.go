package competencies

import (
	"strconv"

	"github.com/pkg/errors"
)

var ErrNoSuchValue = errors.New("competencies type: no such value")

type Type uint8

const (
	GENERAL Type = iota + 1
	PROFESSIONAL
	UNIVERSAL
)

func (t Type) String() string {
	switch t {
	case GENERAL:
		return "ОПК"
	case PROFESSIONAL:
		return "ПК"
	case UNIVERSAL:
		return "УК"
	}

	return "%!CompetencyType(" + strconv.Itoa(int(t)) + ")"
}

func (t Type) IsValid() bool {
	switch t {
	case GENERAL, PROFESSIONAL, UNIVERSAL:
		return true
	}

	return false
}
