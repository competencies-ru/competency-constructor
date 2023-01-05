package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
)

func decodeCreateLevelCommand(w http.ResponseWriter, r *http.Request) (command.CreateLevel, bool) {
	var request CreateLevelRequest
	if ok := decode(w, r, &request); !ok {
		return command.CreateLevel{}, ok
	}

	return command.CreateLevel{Title: request.Title}, true
}
