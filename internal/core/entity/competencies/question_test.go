package competencies_test

import (
	"sort"
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testCases []struct {
	Name        string
	Params      competencies.QuestionParams
	ShouldBeErr bool
	ExpectedErr error
}

// testing negative cases with fields
func TestNewQuestionErrField(t *testing.T) {
	t.Parallel()

	cases := testCases{
		{
			Name: "err_empty_id",
			Params: competencies.QuestionParams{
				ID:             "   ",
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(2),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionIDEmpty,
		},
		{
			Name: "err_description_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "   ",
				QuestionType:   competencies.QuestionType(2),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionDescriptionEmpty,
		},
		{
			Name: "err_rank_segment",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(2),
				Rank:           15,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionRankSegment,
		},
		{
			Name: "err_question_type_invalid",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(20),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionTypeInvalid,
		},

		{
			Name: "err_indicator_id_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(2),
				Rank:           1,
				IndicatorID:    "   ",
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionIndicatorIDEmpty,
		},
	}

	runQuestionTestCases(t, cases)
}

func TestNewQuestionConformity(t *testing.T) {
	cases := testCases{
		{
			Name: "without_error_question_conformity",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(3),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: nil,
				TaskConformity: []competencies.TaskConformityParams{
					{
						Left:  "a",
						Right: "b",
					},
					{
						Left:  "c",
						Right: "b",
					},
				},
				CompleteAnswer: "",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_question_conformity_has_one_element",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(3),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: nil,
				TaskConformity: []competencies.TaskConformityParams{
					{
						Left:  "a",
						Right: "b",
					},
				},
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionTaskConformityIsEmpty,
		},
		{
			Name: "error_question_conformity_is_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(3),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionTaskConformityIsEmpty,
		},
		{
			Name: "error_question_conformity_left_is_empty",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(3),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: nil,
				TaskConformity: []competencies.TaskConformityParams{
					{
						Left: "		",
						Right: "b",
					},
					{
						Left:  "a",
						Right: "b",
					},
				},
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskConformityLeftEmpty,
		},
		{
			Name: "error_question_conformity_right_is_empty",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(3),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: nil,
				TaskConformity: []competencies.TaskConformityParams{
					{
						Right: "		",
						Left: "b",
					},
					{
						Left:  "a",
						Right: "b",
					},
				},
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskConformityRightEmpty,
		},
	}

	runQuestionTestCases(t, cases)
}

// testing complete answer task
func TestNewQuestionCompleteAnswer(t *testing.T) {
	cases := testCases{
		{
			Name: "without_error_question_complete",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(4),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "answer",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_test_sequence_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(2),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   []competencies.TaskSequenceParams{},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionTaskSequenceIsEmpty,
		},
		{
			Name: "error_question_complete_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(4),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				CompleteAnswer: "		",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionCompleteAnswerIsEmpty,
		},
	}

	runQuestionTestCases(t, cases)
}

// test questions sequence
func TestNewQuestionSequence(t *testing.T) {
	cases := testCases{
		{
			Name: "without_error_question_sequence",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Sequence: 1,
						Answer:   "a",
					},
					{
						Sequence: 2,
						Answer:   "b",
					},
					{
						Sequence: 3,
						Answer:   "c",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_test_sequence_has_one_element",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Sequence: 1,
						Answer:   "a",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrQuestionTaskSequenceIsEmpty,
		},
		{
			Name: "error_test_sequence_incorrect",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Sequence: 1,
						Answer:   "a",
					},
					{
						Sequence: 2,
						Answer:   "a",
					},
					{
						Sequence: 4,
						Answer:   "a",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceIncorrect,
		},
		{
			Name: "error_task_sequence_empty_answer",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Sequence: 1,
						Answer:   "",
					},
					{
						Sequence: 2,
						Answer:   "a",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceAnswer,
		},
		{
			Name: "error_task_sequence_empty_sequence",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Answer: "a",
					},
					{
						Sequence: 1,
						Answer:   "a",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceIncorrect,
		},
		{
			Name: "error_task_sequence_incorrect_sequence",
			Params: competencies.QuestionParams{
				ID:           uuid.NewString(),
				Description:  "new questions",
				QuestionType: competencies.QuestionType(2),
				Rank:         1,
				IndicatorID:  uuid.NewString(),
				TaskSequence: []competencies.TaskSequenceParams{
					{
						Sequence: competencies.MaxCountSequence + 1,
						Answer:   "a",
					},
					{
						Sequence: competencies.MaxCountSequence + 2,
						Answer:   "a",
					},
				},
				TaskConformity: nil,
				CompleteAnswer: "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskSequenceIncorrect,
		},
	}

	runQuestionTestCases(t, cases[4:])
}

// testing task point questions
func TestNewQuestionTaskPoints(t *testing.T) {
	cases := testCases{
		{
			Name: "without_error_question_task_point_single",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{1},
				},
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error_question_task_point_not_single",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   false,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{0, 2},
				},
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_task_point_empty_variants",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: nil,
					Answers:  []int{1},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsIsEmpty,
		},
		{
			Name: "error_task_point_max_more",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: make([]string, competencies.MaxVariantsLen+1),
					Answers:  []int{1},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsMoreMax,
		},
		{
			Name: "error_task_variants_is_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: make([]string, competencies.MaxVariantsLen),
					Answers:  []int{1},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointVariantsIsEmptyValue,
		},

		{
			Name: "error_task_answer_is_empty_single",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIsEmpty,
		},
		{
			Name: "error_task_many_answer_is_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   false,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIsEmpty,
		},
		{
			Name: "error_task_single_answer_is_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{1, 2},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_single_answer_is_empty",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   true,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{10},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_one_answer",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   false,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{1},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_more_max_answer",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   false,
					Variants: []string{"a", "b", "c"},
					Answers:  make([]int, competencies.MaxAnswerNumber+1),
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerNumberIncorrect,
		},
		{
			Name: "error_task_many_answer_contains_more_max_answer",
			Params: competencies.QuestionParams{
				ID:             uuid.NewString(),
				Description:    "new questions",
				QuestionType:   competencies.QuestionType(1),
				Rank:           1,
				IndicatorID:    uuid.NewString(),
				TaskSequence:   nil,
				TaskConformity: nil,
				TaskPoint: competencies.TaskPointParams{
					Single:   false,
					Variants: []string{"a", "b", "c"},
					Answers:  []int{-1, 10},
				},
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrTaskPointAnswerIncorrect,
		},
	}

	runQuestionTestCases(t, cases)
}

func runQuestionTestCases(t *testing.T, cases testCases) {
	t.Helper()

	for i := range cases {
		c := cases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			q, err := competencies.NewQuestion(c.Params)

			if c.ShouldBeErr {
				t.Run("err_is", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("id", func(t *testing.T) {
					require.Equal(t, c.Params.ID, q.ID())
				})

				t.Run("description", func(t *testing.T) {
					require.Equal(t, c.Params.Description, q.Description())
				})

				t.Run("rank", func(t *testing.T) {
					require.Equal(t, c.Params.Rank, q.Rank())
				})

				t.Run("indicatorID", func(t *testing.T) {
					require.Equal(t, c.Params.IndicatorID, q.IndicatorID())
				})

				t.Run("type", func(t *testing.T) {
					require.Equal(t, c.Params.QuestionType, q.QuestionType())
				})

				t.Run("taskSequence", func(t *testing.T) {
					equalTestSequence(t, c.Params.TaskSequence, q.TaskSequence())
				})

				t.Run("taskComplete", func(t *testing.T) {
					require.Equal(t, c.Params.CompleteAnswer, q.CompleteAnswer())
				})

				t.Run("taskConformity", func(t *testing.T) {
					equalTestConformity(t, c.Params.TaskConformity, q.TaskConformity())
				})

				t.Run("taskPoint", func(t *testing.T) {
					equalTaskPoint(t, c.Params.TaskPoint, q.TaskPoint())
				})
			})
		})
	}
}

func equalTestSequence(
	t *testing.T,
	params []competencies.TaskSequenceParams,
	result []competencies.TaskSequence,
) {
	t.Helper()

	require.True(t, len(params) == len(result))

	sort.Slice(params, func(i, j int) bool {
		return params[i].Sequence <= params[j].Sequence
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Sequence() <= result[j].Sequence()
	})

	for i := 0; i < len(params); i++ {
		require.Equal(t, params[i].Sequence, result[i].Sequence())
		require.Equal(t, params[i].Answer, result[i].Answer())
	}
}

func equalTestConformity(
	t *testing.T,
	params []competencies.TaskConformityParams,
	result []competencies.TaskConformity,
) {
	t.Helper()

	require.True(t, len(params) == len(result))

	sort.Slice(params, func(i, j int) bool {
		return params[i].Left <= params[j].Left
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Left() <= result[j].Left()
	})

	for i := 0; i < len(params); i++ {
		require.Equal(t, params[i].Left, result[i].Left())
		require.Equal(t, params[i].Right, result[i].Right())
	}
}

func equalTaskPoint(
	t *testing.T,
	params competencies.TaskPointParams,
	result competencies.TaskPoint,
) {
	t.Helper()

	require.Equal(t, params.Single, result.IsSingle())
	require.Equal(t, params.Variants, result.Variants())
	require.Equal(t, params.Answers, result.Answer())
}
