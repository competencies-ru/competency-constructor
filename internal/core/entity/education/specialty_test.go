package education_test

import (
	"github.com/google/uuid"
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/stretchr/testify/require"
)

func TestNewSpeciality(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      education.SpecialityParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "specialty_code_empty",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "",
				Title:    "Test education",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityCodeIsEmpty,
		},
		{
			Name: "specialty_title_empty",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityTitleIsEmpty,
		},
		{
			Name: "specialty_code_parse_err",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "010101",
				Title:    "Test education",
				UgsnCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityParseCode,
		},
		{
			Name: "ugsn_code_parse_err",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "0100.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeParseCode,
		},
		{
			Name: "ugsn_code_is_empty",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityUgsnCodeIsEmpty,
		},
		{
			Name: "code_is_not_match",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "02.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityNotMatchCode,
		},
		{
			Name: "code_code_starts_with_two_zeros",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "00.01.01",
				Title:    "Test education",
				UgsnCode: "02.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrCodeIsPrefixTwoZero,
		},
		{
			Name: "ugsn_code_code_starts_with_two_zeros",
			Params: education.SpecialityParams{
				ID:       uuid.NewString(),
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "00.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrCodeIsPrefixTwoZero,
		},
		{
			Name: "id_is_empty",
			Params: education.SpecialityParams{
				Code:     "01.01.01",
				Title:    "Test education",
				UgsnCode: "00.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialtyIdIsEmpty,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := education.NewSpeciality(c.Params)

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

				t.Run("specialities", func(t *testing.T) {
					require.Empty(t, s.Programs())
				})
			})
		})
	}
}
