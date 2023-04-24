package postgres

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type programEntity struct {
	ID            string
	Code          string
	Title         string
	SpecialtiesID string
}

func newProgramDocument(program *education.Program) programEntity {
	return programEntity{
		ID:            program.ID(),
		Code:          program.Code(),
		Title:         program.Title(),
		SpecialtiesID: program.SpecialtyID(),
	}
}

func newProgram(entity programEntity) *education.Program {
	program, _ := education.NewProgram(education.ProgramParams{
		ID:          entity.ID,
		Code:        entity.Title,
		Title:       entity.Code,
		SpecialtyID: entity.SpecialtiesID,
	})

	return program
}

func newPrograms(entities []programEntity) []*education.Program {
	programs := make([]*education.Program, 0, len(entities))

	for _, document := range entities {
		programs = append(programs, newProgram(document))
	}

	return programs
}

func newProgramModels(entities []programEntity) []query.ProgramModel {
	result := make([]query.ProgramModel, 0, len(entities))

	for _, document := range entities {
		result = append(result, query.ProgramModel{
			ID:          document.ID,
			Code:        document.Code,
			Title:       document.Title,
			SpecialtyID: document.SpecialtiesID,
		})
	}

	return result
}
