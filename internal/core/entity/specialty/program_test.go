package specialty_test

import (
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewProgram(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      specialty.ProgramParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: specialty.ProgramParams{
				ID:            uuid.NewString(),
				Title:         "Test Program",
				SpecialtyCode: "01.10.10",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: specialty.ProgramParams{
				ID:            "",
				Title:         "Test Program",
				SpecialtyCode: "01.10.10",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrProgramIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: specialty.ProgramParams{
				ID:            uuid.NewString(),
				Title:         "",
				SpecialtyCode: "01.10.10",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrProgramTitleIsEmpty,
		},
		{
			Name: "specialty_code_is_empty",
			Params: specialty.ProgramParams{
				ID:            uuid.NewString(),
				Title:         "test",
				SpecialtyCode: "",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrProgramSpecialtyCodeIsEmpty,
		},
		{
			Name: "specialty_code_parse_error",
			Params: specialty.ProgramParams{
				ID:            uuid.NewString(),
				Title:         "test",
				SpecialtyCode: "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrSpecialityParseCode,
		},
		{
			Name: "specialty_code_is_two_zero",
			Params: specialty.ProgramParams{
				ID:            uuid.NewString(),
				Title:         "test",
				SpecialtyCode: "00.10.00",
			},
			ShouldBeErr: true,
			ExpectedErr: specialty.ErrCodeIsPrefixTwoZero,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := specialty.NewProgram(c.Params)

			if c.ShouldBeErr {
				t.Run("err_is", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("code", func(t *testing.T) {
					require.Equal(t, c.Params.ID, s.ID())
				})

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, s.Title())
				})

				t.Run("ugsnCode", func(t *testing.T) {
					require.Equal(t, c.Params.SpecialtyCode, s.SpecialityCode())
				})
			})
		})
	}
}
