package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
)

func decodeCreateLevelCommand(w http.ResponseWriter, r *http.Request) (command.CreateLevel, bool) {
	var request CreateLevelRequest
	if ok := decode(w, r, &request); !ok {
		return command.CreateLevel{}, ok
	}

	return command.CreateLevel{Title: request.Title}, true
}

func renderLevelResponse(w http.ResponseWriter, r *http.Request, models []query.LevelModel) {
	response := make([]LevelResponse, 0, len(models))

	for _, model := range models {
		response = append(response, LevelResponse{
			Id:    toUUID(model.ID),
			Title: model.Title,
		})
	}

	render.Respond(w, r, response)
}
