package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCompetency(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Params      competencies.CompetencyParam
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error_universal",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "УК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(1),
				LevelID:        uuid.NewString(),
				UgsnID:         "",
				SpecialtyID:    "",
				ProgramID:      "",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error_professional",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(3),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    "",
				ProgramID:      uuid.NewString(),
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error_general_for_ugsn",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ОПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         uuid.NewString(),
				SpecialtyID:    "",
				ProgramID:      "",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "without_error_general_for_specialty",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ОПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "error_title_is_empty",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Code:           "ОПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrCompetencyTitleIsEmpty,
		},
		{
			Name: "error_code_is_empty",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrParseCompetencyCode,
		},
		{
			Name: "error_parse_code_universal",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "УК1",
				Category:       "new category",
				CompetencyType: competencies.Type(1),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrParseCompetencyCode,
		},
		{
			Name: "error_parse_code_general",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ОПК",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrParseCompetencyCode,
		},
		{
			Name: "error_parse_code_professional",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК.1",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrParseCompetencyCode,
		},
		{
			Name: "error_invalid_type",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(20),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrCompetencyTypeInvalid,
		},
		{
			Name: "error_invalid_type",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(20),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrCompetencyTypeInvalid,
		},
		{
			Name: "error_invalid_type",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(20),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrCompetencyTypeInvalid,
		},
		{
			Name: "error_education_validate_param_professional",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(3),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    uuid.NewString(),
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrProgramIDIsEmpty,
		},
		{
			Name: "error_education_validate_param_universal",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "УК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(1),
				LevelID:        "",
				UgsnID:         uuid.NewString(),
				SpecialtyID:    "",
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrLevelIDIsEmpty,
		},
		{
			Name: "error_education_validate_param_general",
			Params: competencies.CompetencyParam{
				ID:             uuid.NewString(),
				Title:          "new title",
				Code:           "ОПК-1",
				Category:       "new category",
				CompetencyType: competencies.Type(2),
				LevelID:        "",
				UgsnID:         "",
				SpecialtyID:    "",
				ProgramID:      "",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSpecialtyIDAndUgsnIDIsEmpty,
		},
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			cmp, err := competencies.NewCompetency(c.Params)

			if c.ShouldBeErr {
				t.Run("is_err", func(t *testing.T) {
					require.ErrorIs(t, c.ExpectedErr, err)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("id", func(t *testing.T) {
					require.Equal(t, c.Params.ID, cmp.ID())
				})

				t.Run("title", func(t *testing.T) {
					require.Equal(t, c.Params.Title, cmp.Title())
				})

				t.Run("code", func(t *testing.T) {
					require.Equal(t, c.Params.Code, cmp.Code())
				})

				t.Run("type", func(t *testing.T) {
					require.Equal(t, c.Params.CompetencyType, cmp.CompetencyType())
				})

				t.Run("category", func(t *testing.T) {
					require.Equal(t, c.Params.Category, cmp.Category())

				})

				t.Run("levelID", func(t *testing.T) {
					require.Equal(t, c.Params.LevelID, cmp.LevelID())
				})

				t.Run("ugsnID", func(t *testing.T) {
					require.Equal(t, c.Params.UgsnID, cmp.UgsnID())
				})

				t.Run("specialtyID", func(t *testing.T) {
					require.Equal(t, c.Params.SpecialtyID, cmp.SpecialtyID())

				})

				t.Run("programID", func(t *testing.T) {
					require.Equal(t, c.Params.ProgramID, cmp.ProgramID())

				})
			})

		})
	}
}
