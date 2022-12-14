package specialty_test

import (
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/stretchr/testify/require"
)

func TestNewSpeciality(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      specialty.SpecialityParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test specialty",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "specialty_code_empty",
			Params: specialty.SpecialityParams{
				Code:     "",
				Title:    "Test specialty",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityCodeIsEmpty,
		},
		{
			Name: "specialty_title_empty",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityTitleIsEmpty,
		},
		{
			Name: "specialty_code_parse_err",
			Params: specialty.SpecialityParams{
				Code:     "010101",
				Title:    "Test specialty",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityParseCode,
		},
		{
			Name: "ugsn_code_parse_err",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test specialty",
				UgsnCode: "0100.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrUgsnCodeParseCode,
		},
		{
			Name: "ugsn_code_is_empty",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test specialty",
				UgsnCode: "",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityUgsnCodeIsEmpty,
		},
		{
			Name: "code_is_not_match",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test specialty",
				UgsnCode: "02.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityNotMatchCode,
		},
		{
			Name: "code_code_starts_with_two_zeros",
			Params: specialty.SpecialityParams{
				Code:     "00.01.01",
				Title:    "Test specialty",
				UgsnCode: "02.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrCodeIsPrefixTwoZero,
		},
		{
			Name: "ugsn_code_code_starts_with_two_zeros",
			Params: specialty.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test specialty",
				UgsnCode: "00.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrCodeIsPrefixTwoZero,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := specialty.NewSpeciality(c.Params)

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

				t.Run("ugsnCode", func(t *testing.T) {
					require.Equal(t, c.Params.UgsnCode, s.UgsnCode())
				})

				t.Run("specialities", func(t *testing.T) {
					require.Empty(t, s.Programs())
				})
			})
		})
	}
}
