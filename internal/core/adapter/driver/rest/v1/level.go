package v1

import (
	"fmt"
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
)

func (h handler) GetLevels(w http.ResponseWriter, r *http.Request) {
	models, err := h.app.Queries.FindLevels.Handle(r.Context())
	if err == nil {
		renderLevelResponse(w, r, models)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) CreateLevel(w http.ResponseWriter, r *http.Request) {
	cmd, ok := decodeCreateLevelCommand(w, r)

	if !ok {
		return
	}

	id, err := h.app.Commands.CreateLevel.Handle(r.Context(), cmd)

	if err == nil {
		w.Header().Set(
			"Location",
			fmt.Sprintf("/levels/%s", id),
		)
		w.WriteHeader(http.StatusCreated)

		return
	}

	if errors.Is(err, service.ErrLevelAlreadyExists) {
		rest.NotFound(string(BadRequest), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) GetSpecificLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	// TODO implement me
	panic("implement me")
}
