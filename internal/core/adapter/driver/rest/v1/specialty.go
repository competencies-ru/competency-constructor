package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
)

func (h handler) GetSpecialties(w http.ResponseWriter, r *http.Request, id string) {
	models, err := h.app.Queries.FindAllSpecialties.Handle(r.Context(), id)

	if err == nil {
		renderSpecialtyResponse(w, r, models)

		return
	}

	if errors.Is(err, service.ErrUgsnNotFound) {
		rest.NotFound(string(UgsnNotFound), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) AddSpecialties(w http.ResponseWriter, r *http.Request, id string) {
	command, ok := decodeSpecialtyRequest(w, r)

	if !ok {
		return
	}

	err := h.app.Commands.AddSpecialties.Handle(r.Context(), id, command)

	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	if errors.Is(err, service.ErrUgsnNotFound) {
		rest.NotFound(string(UgsnNotFound), err, w, r)

		return
	}

	if errors.Is(err, service.ErrSpecialtyAlreadyExists) {
		rest.NotFound(string(BadRequest), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) DeleteSpecialty(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetSpecificSpecialty(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	panic("implement me")
}
