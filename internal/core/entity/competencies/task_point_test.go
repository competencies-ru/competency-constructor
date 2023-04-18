package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTaskPoint(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Params      competencies.TaskPointParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error_task_point_single",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{2},
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error_task_point_more",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{0, 2},
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_task_point_empty_variants",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: nil,
				Answers:  []int{1},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsIsEmpty,
		},
		{
			Name: "error_task_point_max_more",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: make([]string, competencies.MaxVariantsLen+1),
				Answers:  []int{1},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsMoreMax,
		},
		{
			Name: "error_task_variants_is_empty",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: make([]string, competencies.MaxVariantsLen),
				Answers:  []int{1},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsIsEmptyValue,
		},
		{
			Name: "error_task_answer_is_empty_single",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIsEmpty,
		},
		{
			Name: "error_task_answer_is_empty_single",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIsEmpty,
		},
		{
			Name: "error_task_many_answer_is_empty",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIsEmpty,
		},
		{
			Name: "error_task_single_answer_is_empty",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{1, 2},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_single_answer_is_empty",
			Params: competencies.TaskPointParams{
				Single:   true,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{10},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_one_answer",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{1},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_more_max_answer",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  make([]int, competencies.MaxAnswerNumber+1),
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_more_max_answer",
			Params: competencies.TaskPointParams{
				Single:   false,
				Variants: []string{"a", "b", "c"},
				Answers:  []int{-1, 10},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIncorrect,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			point, err := competencies.NewTaskPoint(c.Params)

			if c.ShouldBeErr {

				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, c.ExpectedErr, err)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.Equal(t, c.Params.Single, point.IsSingle())
				require.Equal(t, c.Params.Variants, point.Variants())
				require.Equal(t, c.Params.Answers, point.Answer())
			})
		})
	}
}
