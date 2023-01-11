package query

import (
	"context"

	"github.com/pkg/errors"
)

type FilterCompetencyReadModels interface {
	FilterCompetencies(
		ctx context.Context,
		levelID,
		ugsnID,
		specialtyID,
		programID string,
	) ([]CompetencyModel, error)
}

type FilterCompetenciesHandler interface {
	Handle(ctx context.Context, params FilterCompetencyParam) ([]CompetencyModel, error)
}

type filterCompetenciesHandler struct {
	repository FilterCompetencyReadModels
}

func NewFilterCompetenciesHandler(repository FilterCompetencyReadModels) FilterCompetenciesHandler {
	if repository == nil {
		panic("competency repository is nil ")
	}

	return filterCompetenciesHandler{repository: repository}
}

func (f filterCompetenciesHandler) Handle(ctx context.Context, params FilterCompetencyParam) ([]CompetencyModel, error) {
	competencies, err := f.repository.FilterCompetencies(
		ctx,
		params.LevelID,
		params.UgsnID,
		params.SpecialtyID,
		params.ProgramID,
	)

	return competencies, errors.Wrap(err, "filter competency")
}
