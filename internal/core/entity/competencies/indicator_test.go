package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIndicator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Params      competencies.IndicatorParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: competencies.IndicatorParams{
				ID:           uuid.NewString(),
				Title:        "new title",
				Code:         "УК-1.1",
				SubjectID:    uuid.NewString(),
				CompetencyID: uuid.NewString(),
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_id_is_empty",
			Params: competencies.IndicatorParams{
				Title:        "new title",
				Code:         "ПК-1.1",
				SubjectID:    uuid.NewString(),
				CompetencyID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrIndicatorIDIsEmpty,
		},
		{
			Name: "error_title_is_empty",
			Params: competencies.IndicatorParams{
				ID:           uuid.NewString(),
				Code:         "ОПК-1.1",
				SubjectID:    uuid.NewString(),
				CompetencyID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrIndicatorTitleIsEmpty,
		},
		{
			Name: "error_code_invalid",
			Params: competencies.IndicatorParams{
				ID:           uuid.NewString(),
				Title:        "new title",
				Code:         "УПК-11",
				SubjectID:    uuid.NewString(),
				CompetencyID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrParseIndicatorCode,
		},
		{
			Name: "error_competency_id_is_empty",
			Params: competencies.IndicatorParams{
				ID:        uuid.NewString(),
				Title:     "new title",
				Code:      "УК-1.1",
				SubjectID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrIndicatorCompetencyIDIsEmpty,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			indicator, err := competencies.NewIndicator(c.Params)

			if c.ShouldBeErr {
				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, c.ExpectedErr, err)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("id", func(t *testing.T) {
					require.Equal(t, c.Params.ID, indicator.ID())
				})

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, indicator.Title())
				})

				t.Run("code", func(t *testing.T) {
					require.Equal(t, c.Params.Code, indicator.Code())
				})

				t.Run("subjectId", func(t *testing.T) {
					require.Equal(t, c.Params.SubjectID, indicator.SubjectID())
				})

				t.Run("competencyId", func(t *testing.T) {
					require.Equal(t, c.Params.CompetencyID, indicator.CompetencyID())
				})

			})

		})
	}
}

func TestAddSubject(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Name           string
		IndicatorParam competencies.IndicatorParams
		SubjectIDParam string
		ShouldBeErr    bool
		ExpectedErr    error
	}{
		{
			Name:           "without_error",
			SubjectIDParam: uuid.NewString(),
			ShouldBeErr:    false,
			ExpectedErr:    nil,
		},
		{
			Name:        "error_subject_id_is_empty",
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrIndicatorSubjectIDIsEmpty,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			indicator, _ := competencies.NewIndicator(competencies.IndicatorParams{
				ID:           uuid.NewString(),
				Title:        "new title",
				Code:         "УК-1.1",
				SubjectID:    uuid.NewString(),
				CompetencyID: uuid.NewString(),
			})

			err := indicator.AddSubjectID(c.SubjectIDParam)

			if c.ShouldBeErr {
				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, c.ExpectedErr, err)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("subjectID", func(t *testing.T) {
					require.Equal(t, c.SubjectIDParam, indicator.SubjectID())
				})

			})
		})
	}

}
