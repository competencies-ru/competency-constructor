package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"
)

func decodeCreateLevelCommand(w http.ResponseWriter, r *http.Request) (command.CreateLevel, bool) {
	var request CreateLevelRequest
	if ok := decode(w, r, &request); !ok {
		return command.CreateLevel{}, ok
	}

	return command.CreateLevel{Title: request.Title}, true
}

func decodeCreateUgsnCommands(w http.ResponseWriter, r *http.Request) ([]command.CreateUgsnCommand, bool) {
	var requests []CreateUgsnRequest

	if ok := decode(w, r, &requests); !ok {
		return []command.CreateUgsnCommand{}, ok
	}

	commands := make([]command.CreateUgsnCommand, 0, len(requests))
	for _, request := range requests {
		commands = append(commands, command.CreateUgsnCommand{
			Code:  request.Code,
			Title: request.Title,
		})
	}

	return commands, true
}

func decodeCreateSpecialtiesCommands(w http.ResponseWriter, r *http.Request) ([]command.CreateSpecialtiesCommand, bool) {
	var requests []CreateSpecialtyRequest

	if ok := decode(w, r, &requests); !ok {
		return []command.CreateSpecialtiesCommand{}, ok
	}

	commands := make([]command.CreateSpecialtiesCommand, 0, len(requests))
	for _, request := range requests {
		commands = append(commands, command.CreateSpecialtiesCommand{
			Code:  request.Code,
			Title: request.Title,
		})
	}

	return commands, true
}

func decodeCreateProgramCommands(w http.ResponseWriter, r *http.Request) ([]command.CreateProgramCommand, bool) {
	var requests []CreateSpecialtyRequest

	if ok := decode(w, r, &requests); !ok {
		return []command.CreateProgramCommand{}, ok
	}

	commands := make([]command.CreateProgramCommand, 0, len(requests))
	for _, request := range requests {
		commands = append(commands, command.CreateProgramCommand{
			Code:  request.Code,
			Title: request.Title,
		})
	}

	return commands, true
}

func renderLevelResponse(w http.ResponseWriter, r *http.Request, levels []query.LevelModel) {
	response := make([]LevelResponse, 0, len(levels))
	for _, l := range levels {
		response = append(response, LevelResponse{
			Id:    l.ID,
			Title: l.Title,
		})
	}

	render.Respond(w, r, response)
}

func renderSpecificLevelResponse(w http.ResponseWriter, r *http.Request, model query.SpecificLevelModel) {
	response := SpecificLevelResponse{
		Id:    model.ID,
		Title: model.Title,
		Ugsn:  renderSpecificUgsnResponse(model.Ugsn...),
	}

	render.Respond(w, r, response)
}

func renderSpecificUgsnResponse(models ...query.SpecificUgsnModel) []SpecificUgsnResponse {
	result := make([]SpecificUgsnResponse, 0, len(models))

	for _, model := range models {
		result = append(result, SpecificUgsnResponse{
			Id:        model.ID,
			Code:      model.Code,
			Specialty: renderSpecificSpecialtyResponse(model.Specialties...),
			Title:     model.Title,
		})
	}

	return result
}

func renderSpecificSpecialtyResponse(models ...query.SpecificSpecialtyModel) []SpecificSpecialtyResponse {
	result := make([]SpecificSpecialtyResponse, 0, len(models))

	for _, model := range models {
		result = append(result, SpecificSpecialtyResponse{
			Id:      model.ID,
			Code:    model.Code,
			Program: renderSpecificProgramResponse(model.Programs...),
			Title:   model.Title,
		})
	}

	return result
}

func renderSpecificProgramResponse(models ...query.ProgramModel) []ProgramResponse {
	result := make([]ProgramResponse, 0, len(models))

	for _, model := range models {
		result = append(result, ProgramResponse{
			Id:    model.ID,
			Code:  model.Code,
			Title: model.Title,
		})
	}

	return result
}
