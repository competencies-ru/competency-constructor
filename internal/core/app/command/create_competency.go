package command

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateCompetencyCommand struct {
	Code           string
	Title          string
	Category       string
	CompetencyType competencies.Type
	LevelID        string
	UgsnID         string
	SpecialtyID    string
	ProgramID      string
}

type CreateCompetenceHandler interface {
	Handle(ctx context.Context, cmd CreateCompetencyCommand) (string, error)
}

type createCompetencyHandler struct {
	repository  service.CompetencyRepository
	ugsnRepo    service.UgsnRepository
	levelRepo   service.LevelRepository
	specRepo    service.SpecialtyRepository
	programRepo service.ProgramRepository
}

func NewCreateCompetencyHandler(
	repository service.CompetencyRepository,
	levelRepo service.LevelRepository,
	ugsnRepo service.UgsnRepository,
	specRepo service.SpecialtyRepository,
	programRepo service.ProgramRepository,
) CreateCompetenceHandler {
	if repository == nil {
		panic("competencies repository is nil")
	}

	return createCompetencyHandler{
		repository:  repository,
		ugsnRepo:    ugsnRepo,
		levelRepo:   levelRepo,
		specRepo:    specRepo,
		programRepo: programRepo,
	}
}

func (h createCompetencyHandler) Handle(ctx context.Context, cmd CreateCompetencyCommand) (string, error) {
	id := uuid.NewString()

	competency, err := competencies.NewCompetency(competencies.CompetencyParam{
		ID:             id,
		Title:          cmd.Title,
		Code:           cmd.Code,
		Category:       cmd.Category,
		CompetencyType: cmd.CompetencyType,
		LevelID:        cmd.LevelID,
		UgsnID:         cmd.UgsnID,
		SpecialtyID:    cmd.SpecialtyID,
		ProgramID:      cmd.ProgramID,
	})
	if err != nil {
		return "", err
	}

	if err := h.existEducationsRecord(ctx, competency); err != nil {
		return "", err
	}

	return id, h.repository.AddCompetency(ctx, competency)
}

func (h createCompetencyHandler) existEducationsRecord(ctx context.Context, competency *competencies.Competency) (err error) {
	defer func() {
		err = errors.Wrap(err, "an error occurred while creating the competency")
	}()

	switch competency.CompetencyType() {
	case competencies.UNIVERSAL:
		_, err = h.levelRepo.GetLevel(ctx, competency.LevelID())

		return err
	case competencies.PROFESSIONAL:
		_, err = h.programRepo.GetProgram(ctx, competency.ProgramID())

		return err

	case competencies.GENERAL:
		if competency.UgsnID() != "" {
			_, err = h.ugsnRepo.GetUgsn(ctx, competency.UgsnID())

			return err
		}

		_, err = h.specRepo.GetSpecialty(ctx, competency.SpecialtyID())

		return err
	}

	return
}
