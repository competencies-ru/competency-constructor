package education_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLevel(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      education.LevelParam
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: education.LevelParam{
				ID:    uuid.NewString(),
				Title: "Test Level",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: education.LevelParam{
				Title: "Test Level",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrLevelIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: education.LevelParam{
				ID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrLevelTitleIsEmpty,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := education.NewLevel(c.Params)

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

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, s.Title())
				})

			})
		})
	}
}
