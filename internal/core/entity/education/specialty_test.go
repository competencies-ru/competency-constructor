package education_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
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
				ID:     uuid.NewString(),
				Title:  "Test Specialty",
				UgsnID: uuid.NewString(),
				Code:   "01.01.01",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: education.SpecialityParams{
				Title:  "Test Specialty",
				UgsnID: uuid.NewString(),
				Code:   "01.01.01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialtyIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: education.SpecialityParams{
				ID:     uuid.NewString(),
				UgsnID: uuid.NewString(),
				Code:   "01.01.01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityTitleIsEmpty,
		},
		{
			Name: "code_is_empty",
			Params: education.SpecialityParams{
				ID:     uuid.NewString(),
				Title:  "Test Specialty",
				UgsnID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityCodeIsEmpty,
		},
		{
			Name: "err_parse_code",
			Params: education.SpecialityParams{
				ID:     uuid.NewString(),
				Title:  "Test Specialty",
				Code:   "01.00.00",
				UgsnID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityParseCode,
		},
		{
			Name: "ugsnID_is_empty",
			Params: education.SpecialityParams{
				ID:    uuid.NewString(),
				Title: "Test Specialty",
				Code:  "01.01.01",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrSpecialityUgsnIDIsEmpty,
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
					require.Equal(t, c.Params.UgsnID, s.UgsnID())
				})
			})
		})
	}

}
