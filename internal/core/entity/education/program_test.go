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
				ID:          uuid.NewString(),
				Title:       "Test Program",
				SpecialtyID: uuid.NewString(),
				Code:        "01.01.01-01",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: education.ProgramParams{
				ID:          "",
				Title:       "Test Program",
				SpecialtyID: uuid.NewString(),
				Code:        "01.01.01-01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: education.ProgramParams{
				ID:          uuid.NewString(),
				Title:       "",
				SpecialtyID: uuid.NewString(),
				Code:        "01.01.01-01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramTitleIsEmpty,
		},
		{
			Name: "code_is_empty",
			Params: education.ProgramParams{
				ID:          uuid.NewString(),
				Title:       "title",
				SpecialtyID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramCodeIsEmpty,
		},
		{
			Name: "specialty_id_is_empty",
			Params: education.ProgramParams{
				ID:    uuid.NewString(),
				Title: "title",
				Code:  "10.10.10-01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramSpecialtyIDIsEmpty,
		},
		{
			Name: "err_parse_code",
			Params: education.ProgramParams{
				ID:          uuid.NewString(),
				Title:       "title",
				SpecialtyID: uuid.NewString(),
				Code:        "10.10.10",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrProgramParseCode,
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

				t.Run("id", func(t *testing.T) {
					require.Equal(t, c.Params.ID, s.ID())
				})

				t.Run("code", func(t *testing.T) {
					require.Equal(t, c.Params.Code, s.Code())
				})

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, s.Title())
				})

				t.Run("specialtyID", func(t *testing.T) {
					require.Equal(t, c.Params.SpecialtyID, s.SpecialtyID())
				})
			})
		})
	}
}
