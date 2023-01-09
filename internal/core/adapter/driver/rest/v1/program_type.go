package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"
)

func decodeProgramRequest(w http.ResponseWriter, r *http.Request) (command.CreateProgramCommand, bool) {
	var request CreateProgramRequest

	if ok := decode(w, r, &request); !ok {
		return command.CreateProgramCommand{}, ok
	}

	return command.CreateProgramCommand{Code: request.Code, Title: request.Title}, true
}

func renderProgramsResponse(w http.ResponseWriter, r *http.Request, models []query.ProgramModel) {
	response := make([]ProgramResponse, 0, len(models))

	for _, model := range models {
		response = append(response, newProgramResponse(model))
	}

	render.Respond(w, r, response)
}

func newProgramResponse(model query.ProgramModel) ProgramResponse {
	return ProgramResponse{
		Code:        model.Code,
		Id:          toUUID(model.ID),
		Title:       model.Title,
		SpecialtyId: toUUID(model.SpecialtyID),
	}
}
