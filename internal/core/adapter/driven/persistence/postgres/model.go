package postgres

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"go.uber.org/multierr"
)

type (
	ugsn struct {
		Code  string
		Title string
	}
)

func mapUgsn(input []*ugsn) ([]*specialty.Ugsn, error) {
	var result error

	output := make([]*specialty.Ugsn, 0, len(input))

	for i := range input {
		v := input[i]

		newUgsn, err := specialty.NewUgsn(specialty.UgsnParams{
			Code:  v.Code,
			Title: v.Title,
		})
		if err != nil {
			result = multierr.Append(result, err)

			continue
		}

		output = append(output, newUgsn)
	}

	return output, result
}
