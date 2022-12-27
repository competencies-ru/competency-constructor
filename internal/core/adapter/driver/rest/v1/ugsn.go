package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

func (h handler) GetUgsn(w http.ResponseWriter, r *http.Request, levelID string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) AddUgsn(w http.ResponseWriter, r *http.Request, levelID string) {
	commands, ok := decodeCreateUgsnCommands(w, r)
	if !ok {
		return
	}

	err := h.app.Commands.AddUgsn.Handle(r.Context(), levelID, commands)
	if err == nil {
		return
	}

	if education.IsInvalidUgsnParametersError(err) {
		rest.UnprocessableEntity(string(InvalidJson), err, w, r)

		return
	}

	rest.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) DeleteUgsn(w http.ResponseWriter, r *http.Request, levelID string, ugsnCode string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetSpecificUgsn(w http.ResponseWriter, r *http.Request, levelID string, ugsnCode string) {
	// TODO implement me
	panic("implement me")
}
