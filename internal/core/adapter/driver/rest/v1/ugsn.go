package v1

import (
	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/pkg/errors"
	"net/http"
)

func (h handler) GetUgsn(w http.ResponseWriter, r *http.Request, levelID string) {
	models, err := h.app.Queries.FindAllUgsn.Handle(r.Context(), levelID)
	if err == nil {
		renderUgsnReponse(w, r, models)

		return
	}

	if errors.Is(err, service.ErrLevelNotFound) {
		rest.NotFound(string(LevelNotFound), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) AddUgsn(w http.ResponseWriter, r *http.Request, levelID string) {
	cmd, ok := decodeCreateUgsnResponse(w, r)

	if !ok {
		return
	}

	err := h.app.Commands.AddUgsn.Handle(r.Context(), levelID, cmd)

	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	if errors.Is(err, service.ErrLevelNotFound) {
		rest.NotFound(string(LevelNotFound), err, w, r)

		return
	}

	if errors.Is(err, service.ErrUgsnAlreadyExists) {
		rest.NotFound(string(BadRequest), err, w, r)

		return
	}

	rest.InternalServerError(string(UnexpectedError), err, w, r)
}

func (h handler) DeleteUgsn(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetSpecificUgsn(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}
