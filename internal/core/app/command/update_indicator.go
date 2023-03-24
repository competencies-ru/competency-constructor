package command

import "context"

type UpdateIndicatorCommand struct {
	Title     string
	Code      string
	SubjectID string
}

type UpdateIndicatorHandler interface {
	Handler(ctx context.Context, id string, upd UpdateIndicatorCommand) error
}
