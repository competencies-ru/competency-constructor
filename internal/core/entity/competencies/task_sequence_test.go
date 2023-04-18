package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTaskSequence(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Params      competencies.TaskSequenceParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: competencies.TaskSequenceParams{
				Sequence: 1,
				Answer:   "b",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_empty_sequence",
			Params: competencies.TaskSequenceParams{
				Answer: "b",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceZero,
		},
		{
			Name: "error_incorrect_sequence",
			Params: competencies.TaskSequenceParams{
				Sequence: competencies.MaxCountSequence + 1,
				Answer:   "b",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceIncorrect,
		},
		{
			Name: "error_empty_answer",
			Params: competencies.TaskSequenceParams{
				Sequence: 1,
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceAnswer,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			sequence, err := competencies.NewTaskSequence(c.Params)

			if c.ShouldBeErr {
				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, c.ExpectedErr, err)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("sequence", func(t *testing.T) {
					require.Equal(t, c.Params.Sequence, sequence.Sequence())
				})

				t.Run("answer", func(t *testing.T) {
					require.Equal(t, c.Params.Answer, sequence.Answer())
				})
			})

		})
	}
}
