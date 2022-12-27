package app

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
)

type (
	Application struct {
		Commands Commands
		Queries  Queries
	}

	Commands struct {
		CreateLevel    command.CreateLevelHandler
		AddUgsn        command.AddUgsnHandler
		AddSpecialties command.AddSpecialtiesHandler
		AddPrograms    command.AddProgramHandler
	}

	Queries struct {
		FindLevels        query.FindLevelsHandler
		GetSpecificLevels query.SpecificLevelHandler
	}
)
