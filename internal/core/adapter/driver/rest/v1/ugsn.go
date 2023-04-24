package v1

import (
	"net/http"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
)

func (h handler) GetUgsn(w http.ResponseWriter, r *http.Request, levelID openapiTypes.UUID) {
	models, err := h.app.Queries.FindAllUgsn.Handle(r.Context(), levelID.String())
	if err == nil {
		renderUgsnReponse(w, r, models)

		return
	}

	if errors.Is(err, service.NotFoundEntity) {
		rest.NotFound(string(LevelNotFound), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) AddUgsn(w http.ResponseWriter, r *http.Request, levelID openapiTypes.UUID) {
	cmd, ok := decodeCreateUgsnResponse(w, r)

	if !ok {
		return
	}

	err := h.app.Commands.AddUgsn.Handle(r.Context(), levelID.String(), cmd)

	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	if errors.Is(err, service.NotFoundEntity) {
		rest.NotFound(string(LevelNotFound), err, w, r)

		return
	}

	if errors.Is(err, service.AlreadyExistsEntity) {
		rest.NotFound(string(BadRequest), err, w, r)

		return
	}

	if education.IsInvalidUgsnParametersError(err) {
		rest.UnprocessableEntity(string(InvalidUgsnParameters), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) DeleteUgsn(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetSpecificUgsn(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {
	// TODO implement me
	panic("implement me")
}
