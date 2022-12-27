package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

func (h handler) GetLevels(w http.ResponseWriter, r *http.Request) {
	result, err := h.app.Queries.FindLevels.Handle(r.Context())
	if err == nil {
		renderLevelResponse(w, r, result)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) CreateLevel(w http.ResponseWriter, r *http.Request) {
	command, ok := decodeCreateLevelCommand(w, r)
	if !ok {
		return
	}

	id, err := h.app.Commands.CreateLevel.Handle(r.Context(), command)
	if err == nil {
		w.Header().Set("Location", fmt.Sprintf("/levels/%s", id))
		w.WriteHeader(http.StatusCreated)

		return
	}

	if education.IsInvalidLevelParametersError(err) {
		rest.UnprocessableEntity(string(InvalidLevelParameters), err, w, r)

		return
	}

	rest.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) GetSpecificLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	level, err := h.app.Queries.GetSpecificLevels.Handle(r.Context(), levelID)
	if err == nil {
		renderSpecificLevelResponse(w, r, level)

		return
	}

	if errors.Is(err, service.ErrLevelNotFound) {
		rest.NotFound(string(LevelNotFound), err, w, r)

		return
	}

	rest.InternalServerError("unexpected-error", err, w, r)
}
