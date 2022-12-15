package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
)

func (h handler) GetUgsn(w http.ResponseWriter, r *http.Request) {
	ugsn, err := h.app.Services.UgsnService.GetAllUgsn(r.Context())
	if err != nil {
		rest.BadRequest(string(BadRequest), err, w, r)

		return
	}

	renderUgsnResponse(w, r, ugsn)
}

func (h handler) CreateUgsn(w http.ResponseWriter, r *http.Request) {
	uc, ok := unmarshallingCreateUgsn(w, r)
	if !ok {
		return
	}

	err := h.app.Services.UgsnService.Create(r.Context(), uc)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	if errors.Is(err, service.ErrUgsnAlreadyExists) {
		rest.BadRequest(string(BadRequest), err, w, r)

		return
	}

	if specialty.IsInvalidUgsnParametersError(err) {
		rest.UnprocessableEntity("unsupported-entity", err, w, r)

		return
	}

	rest.InternalServerError("internal-server", err, w, r)
}

func (h handler) GetSpecificUgsn(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	specific, err := h.app.Services.UgsnService.GetSpecificUgsn(r.Context(), ugsnCode)

	if err == nil {
		renderSpecificUgsn(w, r, specific)

		return
	}

	if specialty.IsInvalidUgsnParametersError(err) {
		rest.UnprocessableEntity("unsupported-entity", err, w, r)

		return
	}

	if errors.Is(err, service.ErrUgsnNotFound) {
		rest.NotFound("not-found", err, w, r)

		return
	}

	rest.InternalServerError("internal-server", err, w, r)
}

func (h handler) GetSpecialties(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) AddedSpecialties(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}
