package command

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
)

type CreateSubjectCommand struct {
	Name  string
	SName string
}

type SubjectCreator interface {
	AddSubject(ctx context.Context, subject *competencies.Subject) error
}

type CreateSubjectHandler interface {
	Handle(ctx context.Context, command CreateSubjectCommand) error
}

type createSubjectHandle struct {
	creator SubjectCreator
}

func NewCreateSubjectHandler(creator SubjectCreator) CreateSubjectHandler {
	if creator == nil {
		panic("subject creator is nil")
	}

	return createSubjectHandle{creator: creator}
}

func (c createSubjectHandle) Handle(ctx context.Context, cmd CreateSubjectCommand) error {

	sid := uuid.NewString()

	subject, err := competencies.NewSubject(competencies.SubjectParams{
		ID:    sid,
		Name:  cmd.Name,
		Sname: cmd.SName,
	})

	if err != nil {
		return err
	}

	return c.creator.AddSubject(ctx, subject)
}
