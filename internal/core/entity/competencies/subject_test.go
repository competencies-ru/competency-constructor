package competencies_test

import (
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSubject(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Params      competencies.SubjectParams
		ShouldBeErr bool
		ExpectedErr error
	}{
		{
			Name: "without_error",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Name:  "New subject",
				Sname: "N.Subject",
			},
			ShouldBeErr: false,
			ExpectedErr: nil,
		},
		{
			Name: "id_is_empty",
			Params: competencies.SubjectParams{
				Name:  "New subject",
				Sname: "N.Subject",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSubjectEmptyID,
		},
		{
			Name: "sName_is_empty",
			Params: competencies.SubjectParams{
				ID:   uuid.NewString(),
				Name: "New subject",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSubjectEmptySName,
		},
		{
			Name: "name_is_empty",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Sname: "N.subject",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSubjectEmptyName,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, err := competencies.NewSubject(c.Params)

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

				t.Run("name", func(t *testing.T) {
					require.Equal(t, c.Params.Name, s.Name())
				})

				t.Run("sName", func(t *testing.T) {
					require.Equal(t, c.Params.Sname, s.SName())
				})

			})

		})
	}
}

func TestSubjectChangeName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name           string
		Params         competencies.SubjectParams
		NewNameSubject string
		ShouldBeErr    bool
		ExpectedErr    error
	}{
		{
			Name: "without_error",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Name:  "New subject",
				Sname: "N.Subject",
			},
			NewNameSubject: "Subject Name 2",
			ShouldBeErr:    false,
			ExpectedErr:    nil,
		},

		{
			Name: "name_is_empty",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Name:  "New subject",
				Sname: "N.Subject",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSubjectEmptyName,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, _ := competencies.NewSubject(c.Params)

			err := s.ChangeName(c.NewNameSubject)

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

				t.Run("name", func(t *testing.T) {
					require.Equal(t, c.NewNameSubject, s.Name())
				})

				t.Run("sName", func(t *testing.T) {
					require.Equal(t, c.Params.Sname, s.SName())
				})

			})

		})
	}
}

func TestSubjectChangeSName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name            string
		Params          competencies.SubjectParams
		NewSNameSubject string
		ShouldBeErr     bool
		ExpectedErr     error
	}{
		{
			Name: "without_error",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Name:  "New subject",
				Sname: "N.Subject",
			},
			NewSNameSubject: "Subject Name 2",
			ShouldBeErr:     false,
			ExpectedErr:     nil,
		},

		{
			Name: "sname_is_empty",
			Params: competencies.SubjectParams{
				ID:    uuid.NewString(),
				Name:  "New subject",
				Sname: "N.Subject",
			},
			ShouldBeErr: true,
			ExpectedErr: competencies.ErrSubjectEmptySName,
		},
	}

	for i := range testCases {
		c := testCases[i]

		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()

			s, _ := competencies.NewSubject(c.Params)

			err := s.ChangeSName(c.NewSNameSubject)

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

				t.Run("name", func(t *testing.T) {
					require.Equal(t, c.Params.Name, s.Name())
				})

				t.Run("sName", func(t *testing.T) {
					require.Equal(t, c.NewSNameSubject, s.SName())
				})

			})

		})
	}
}
