package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
)

func (h handler) DeleteProgram(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetPrograms(w http.ResponseWriter, r *http.Request, id string) {
	models, err := h.app.Queries.FindAllPrograms.Handle(r.Context(), id)

	if err == nil {
		renderProgramsResponse(w, r, models)

		return
	}

	if errors.Is(err, service.ErrSpecialtyNotFound) {
		rest.NotFound(string(SpecialtyNotFound), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) AddPrograms(w http.ResponseWriter, r *http.Request, id string) {
	command, ok := decodeProgramRequest(w, r)

	if !ok {
		return
	}

	err := h.app.Commands.AddPrograms.Handle(r.Context(), id, command)

	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	if errors.Is(err, service.ErrSpecialtyNotFound) {
		rest.NotFound(string(SpecialtyNotFound), err, w, r)

		return
	}

	if errors.Is(err, service.ErrProgramAlreadyExists) {
		rest.NotFound(string(BadRequest), err, w, r)

		return
	}

	if education.IsInvalidProgramParameters(err) {
		rest.NotFound(string(InvalidSpecialtiesParameters), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}
