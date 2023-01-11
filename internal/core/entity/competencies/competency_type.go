package competencies

import (
	"github.com/pkg/errors"
	"strconv"
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
		return "OПК"
	case PROFESSIONAL:
		return "ПК"
	case UNIVERSAL:
		return "УК"
	}

	return "%!CompetencyType(" + strconv.Itoa(t) + ")"
}

func (t Type) IsValid() bool {
	switch t {
	case GENERAL, PROFESSIONAL, UNIVERSAL:
		return true
	}

	return false
}
