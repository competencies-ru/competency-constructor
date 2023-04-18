package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTaskConformity(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Params      competencies.TaskConformityParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: competencies.TaskConformityParams{
				Left:  "a",
				Right: "b",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error",
			Params: competencies.TaskConformityParams{
				Right: "b",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskConformityLeftEmpty,
		},
		{
			Name: "without_error",
			Params: competencies.TaskConformityParams{
				Left: "b",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskConformityRightEmpty,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			conformity, err := competencies.NewTaskConformity(c.Params)

			if c.ShouldBeErr {
				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("left", func(t *testing.T) {
					require.Equal(t, conformity.Left(), c.Params.Left)
				})

				t.Run("right", func(t *testing.T) {
					require.Equal(t, conformity.Right(), c.Params.Right)
				})
			})

		})
	}
}
