package v1

import (
	"fmt"
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
)

func (h handler) FilterCompetency(w http.ResponseWriter, r *http.Request, params FilterCompetencyParams) {
	filterCompetencyParams := mappingFilterCompetencyParams(params)

	models, err := h.app.Queries.FindAllCompetency.Handle(r.Context(), filterCompetencyParams)
	if err == nil {
		renderCompetenciesResponse(w, r, models)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) CreateCompetency(w http.ResponseWriter, r *http.Request) {
	cmd, ok := decodeCreateCompetencyRequest(w, r)

	if !ok {
		return
	}

	id, err := h.app.Commands.CreateCompetency.Handle(r.Context(), cmd)
	if err == nil {
		w.Header().Set(
			"Location",
			fmt.Sprintf("/competencies/%s", id),
		)
		w.WriteHeader(http.StatusCreated)

		return
	}

	if errors.Is(err, service.ErrCompetencyAlreadyExists) {
		rest.BadRequest(string(BadRequest), err, w, r)

		return
	}

	if service.IsNotFoundEntity(err) {
		rest.NotFound(string(NotFoundEntity), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}
