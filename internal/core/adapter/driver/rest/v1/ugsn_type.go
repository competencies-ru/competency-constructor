package v1

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"
	"net/http"
)

func renderUgsnReponse(w http.ResponseWriter, r *http.Request, models []query.UgsnModel) {
	response := make([]UgsnResponse, 0, len(models))

	for _, model := range models {
		response = append(response, UgsnResponse{
			Id:      model.ID,
			Code:    model.Code,
			Title:   model.Title,
			LevelId: model.LevelID,
		})
	}

	render.Respond(w, r, response)
}

func decodeCreateUgsnResponse(w http.ResponseWriter, r *http.Request) (command.CreateUgsnCommand, bool) {
	var request CreateUgsnRequest

	if !decode(w, r, &request) {
		return command.CreateUgsnCommand{}, false
	}

	return command.CreateUgsnCommand{
		Code:  request.Code,
		Title: request.Title,
	}, true
}
