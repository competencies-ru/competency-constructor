package education_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringRunes(charset string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[b2.Int64()]
	}

	return string(b)
}

func TestNewUgsn(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      education.UgsnParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: education.UgsnParams{
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnIDIsEmpty,
		},
		{
			Name: "title_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnTitleIsEmpty,
		},
		{
			Name: "levelID_is_empty",
			Params: education.UgsnParams{
				ID:    uuid.NewString(),
				Title: "Test ugsn",
				Code:  "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnLevelIDEmpty,
		},
		{
			Name: "code_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeIsEmpty,
		},
		{
			Name: "code_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeIsEmpty,
		},
		{
			Name: "code_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.0000",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnParseCode,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := education.NewUgsn(c.Params)

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

				t.Run("levelID", func(t *testing.T) {
					require.Equal(t, c.Params.LevelID, s.LeveID())
				})
			})
		})
	}
}

func TestRenameUgsn(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		NewTitle    string
		Params      education.UgsnParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name:     "without_error",
			NewTitle: "New Title",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "title_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnTitleIsEmpty,
		},
		{
			Name: "title_max_len",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			NewTitle:    randStringRunes(charset, education.MaxLenTitle+1),
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnTitleMaxLenTitle,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, _ := education.NewUgsn(c.Params)

			err := s.Rename(c.NewTitle)

			if c.ShouldBeErr {
				t.Run("err_is", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("newTitle", func(t *testing.T) {
					require.Equal(t, c.NewTitle, s.Title())
				})
			})
		})
	}
}

func TestUgsnChangeCode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		NewCode     string
		Params      education.UgsnParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name:    "without_error",
			NewCode: "02.00.00",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "code_is_empty",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnCodeIsEmpty,
		},
		{
			Name: "title_max_len",
			Params: education.UgsnParams{
				ID:      uuid.NewString(),
				Title:   "Test ugsn",
				LevelID: uuid.NewString(),
				Code:    "01.00.00",
			},
			NewCode:     "00.00.00",
			ShouldBeErr: true,
			ExpectedErr: education.ErrUgsnParseCode,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, _ := education.NewUgsn(c.Params)

			err := s.ChangeCode(c.NewCode)

			if c.ShouldBeErr {
				t.Run("err_is", func(t *testing.T) {
					require.ErrorIs(t, err, c.ExpectedErr)
				})

				return
			}

			t.Run("no_err", func(t *testing.T) {
				require.NoError(t, err)

				t.Run("newCode", func(t *testing.T) {
					require.Equal(t, c.NewCode, s.Code())
				})
			})
		})
	}
}
