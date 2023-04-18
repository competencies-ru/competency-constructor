package query

import "context"

type SubjectFilterProvider interface {
	FilterSubjectsByName(ctx context.Context, name string) []SubjectModel
}

type SubjectHandler interface {
	Handle(ctx context.Context, name string) []SubjectModel
}

type subjectHandler struct {
	provider SubjectFilterProvider
}

func NewSubjectHandler(provider SubjectFilterProvider) SubjectHandler {
	if provider == nil {
		panic("subject provider is nil")
	}

	return subjectHandler{provider: provider}
}

func (s subjectHandler) Handle(ctx context.Context, name string) []SubjectModel {
	return s.provider.FilterSubjectsByName(ctx, name)
}
