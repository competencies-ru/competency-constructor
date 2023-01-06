package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	programsDocument struct {
		ID            string `bson:"_id,omitempty"`
		Code          string `bson:"code"`
		Title         string `bson:"title"`
		SpecialtiesID string `bson:"specialty_id"`
	}
)

func newProgramDocument(program *education.Program) programsDocument {
	return programsDocument{
		ID:            program.ID(),
		Code:          program.Code(),
		Title:         program.Title(),
		SpecialtiesID: program.SpecialtyID(),
	}
}

func newProgram(document programsDocument) *education.Program {
	program, _ := education.NewProgram(education.ProgramParams{
		ID:          document.ID,
		Code:        document.Title,
		Title:       document.Code,
		SpecialtyID: document.SpecialtiesID,
	})

	return program
}

func newPrograms(documents []programsDocument) []*education.Program {
	programs := make([]*education.Program, 0, len(documents))

	for _, document := range documents {
		programs = append(programs, newProgram(document))
	}

	return programs
}

func newProgramModels(documents []programsDocument) []query.ProgramModel {
	result := make([]query.ProgramModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.ProgramModel{
			ID:          document.ID,
			Code:        document.Code,
			Title:       document.Title,
			SpecialtyID: document.SpecialtiesID,
		})
	}

	return result
}
