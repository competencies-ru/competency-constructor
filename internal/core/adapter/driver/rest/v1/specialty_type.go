package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"
)

func decodeSpecialtyRequest(w http.ResponseWriter, r *http.Request) (command.CreateSpecialtyCommand, bool) {
	var request CreateSpecialtyRequest

	if ok := decode(w, r, &request); !ok {
		return command.CreateSpecialtyCommand{}, ok
	}

	return command.CreateSpecialtyCommand{Code: request.Code, Title: request.Title}, true
}

func renderSpecialtyResponse(w http.ResponseWriter, r *http.Request, models []query.SpecialtyModel) {
	response := make([]SpecialtyResponse, 0, len(models))

	for _, model := range models {
		response = append(response, SpecialtyResponse{
			Code:   model.Code,
			Id:     model.ID,
			Title:  model.Title,
			UgsnId: model.UgsnID,
		})
	}

	render.Respond(w, r, response)
}
