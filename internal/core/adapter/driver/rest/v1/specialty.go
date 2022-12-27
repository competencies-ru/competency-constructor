package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

func (h handler) DeleteSpecialty(w http.ResponseWriter, r *http.Request, levelID, ugsnCode, specialtyCode string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetSpecificSpecialty(w http.ResponseWriter, r *http.Request, levelID, ugsnCode, specialtyCode string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) GetSpecialties(w http.ResponseWriter, r *http.Request, levelID, ugsnCode string) {
	// TODO implement me
	panic("implement me")
}

func (h handler) AddSpecialties(w http.ResponseWriter, r *http.Request, levelID, ugsnCode string) {
	commands, ok := decodeCreateSpecialtiesCommands(w, r)
	if !ok {
		return
	}

	err := h.app.Commands.AddSpecialties.Handle(r.Context(), levelID, ugsnCode, commands)
	if err == nil {
		return
	}

	if education.IsInvalidSpecialtyParametersError(err) {
		rest.UnprocessableEntity(string(InvalidJson), err, w, r)

		return
	}

	rest.InternalServerError("unexpected-error", err, w, r)
}
