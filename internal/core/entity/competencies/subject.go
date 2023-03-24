package competencies

import "github.com/pkg/errors"

type (
	Subject struct {
		id    string
		name  string
		sname string
	}

	SubjectParams struct {
		ID    string
		Name  string
		Sname string
	}
)

var (
	ErrSubjectEmptyID    = errors.New("Subject ID cannot be empty")
	ErrSubjectEmptyName  = errors.New("Subject name cannot be empty")
	ErrSubjectEmptySName = errors.New("Subject surname cannot be empty")
)

func NewSubject(param SubjectParams) (*Subject, error) {
	if param.ID == "" {
		return nil, ErrSubjectEmptyID
	}

	if param.Name == "" {
		return nil, ErrSubjectEmptyName
	}

	if param.Sname == "" {
		return nil, ErrSubjectEmptySName
	}

	return &Subject{id: param.ID, name: param.Name, sname: param.Sname}, nil
}

func (s *Subject) ID() string {
	return s.id
}

func (s *Subject) Name() string {
	return s.name
}

func (s *Subject) SName() string {
	return s.sname
}
