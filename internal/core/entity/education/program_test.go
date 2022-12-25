package education_test

import (
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewProgram(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      education.ProgramParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "29.40.31-25",
				Title:         "Test Program",
				SpecialtyCode: "29.40.31",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "code_is_empty",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "",
				Title:         "Test Program",
				SpecialtyCode: "01.10.10",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "01.10.10-01",
				Title:         "",
				SpecialtyCode: "01.10.10",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramTitleIsEmpty,
		},
		{
			Name: "specialty_code_is_empty",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "01.10.10-01",
				Title:         "test",
				SpecialtyCode: "",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramSpecialtyCodeIsEmpty,
		},
		{
			Name: "specialty_code_parse_error",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "01.10.10-01",
				Title:         "test",
				SpecialtyCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityParseCode,
		},
		{
			Name: "specialty_code_is_two_zero",
			Params: education.ProgramParams{
				ID:            uuid.NewString(),
				Code:          "01.10.10-01",
				Title:         "test",
				SpecialtyCode: "00.10.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrCodeIsPrefixTwoZero,
		},
		{
			Name: "id_is_empty",
			Params: education.ProgramParams{
				Code:          "01.10.10-01",
				Title:         "test",
				SpecialtyCode: "00.10.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrCodeIsPrefixTwoZero,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := education.NewProgram(c.Params)

			if c.ShouldBeErr {
				t.Run("err_is", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("code", func(t *testing.T) {
					require.Equal(t, c.Params.Code, s.Code())
				})

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, s.Title())
				})
			})
		})
	}
}
