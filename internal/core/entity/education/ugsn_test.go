package education_test

import (
	"github.com/google/uuid"
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/stretchr/testify/require"
)

func TestNewUgsn(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      education.UgsnParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Code:  "01.00.00",
				Title: "Test Ugsn",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "ugsn_code_empty",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Code:  "",
				Title: "Test Ugsn",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeIsEmpty,
		},
		{
			Name: "ugsn_title_empty",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Code:  "01.00.00",
				Title: "",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnTitleIsEmpty,
		},
		{
			Name: "ugsn_code_parse_err",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Code:  "010000",
				Title: "Test Ugsn",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeParseCode,
		},
		{
			Name: "ugsn_code_code_starts_with_two_zeros",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Code:  "00.00.00",
				Title: "Test Ugsn",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrCodeIsPrefixTwoZero,
		},
		{
			Name: "ugsn_code_code_starts_with_two_zeros",
			Params: education.UgsnParams{
				Code:  "00.00.00",
				Title: "Test Ugsn",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnIDIsEmpty,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := education.NewUgsn(c.Params)

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
					require.Empty(t, s.Specialities())
				})
			})
		})
	}
}
