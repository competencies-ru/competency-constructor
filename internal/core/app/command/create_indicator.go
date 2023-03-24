package command

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateIndicatorCommand struct {
	Code         string
	Title        string
	CompetencyID string
	SubjectID    string
}

type CreateIndicatorHandler interface {
	Handle(ctx context.Context, competencyID string, cmd CreateIndicatorCommand) (string, error)
}

type createIndicatorHandler struct {
	repository           service.IndicatorRepository
	competencyRepository service.CompetencyRepository
}

func NewCreateIndicatorHandler(
	repository service.IndicatorRepository,
	competencyRepository service.CompetencyRepository,
) CreateIndicatorHandler {

	if repository == nil {
		panic("indicator repository is nil")
	}

	if competencyRepository == nil {
		panic("competency repository is nil")
	}

	return createIndicatorHandler{
		repository:           repository,
		competencyRepository: competencyRepository,
	}
}

func (h createIndicatorHandler) Handle(
	ctx context.Context,
	competencyID string,
	cmd CreateIndicatorCommand,
) (id string, err error) {

	defer func() {
		err = errors.Wrapf(err, "create indicator by competency id: %s", competencyID)
	}()

	if _, err = h.competencyRepository.GetCompetency(ctx, competencyID); err != nil {
		return
	}

	id = uuid.NewString()

	indicator, err := competencies.NewIndicator(competencies.IndicatorParams{
		ID:           id,
		Title:        cmd.Title,
		Code:         cmd.Code,
		SubjectID:    cmd.SubjectID,
		CompetencyID: cmd.CompetencyID,
	})
	if err != nil {
		return "", err
	}

	return id, h.repository.AddIndicator(ctx, indicator)
}
