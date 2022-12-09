package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/go-chi/render"
)

func renderUgsnResponse(
	w http.ResponseWriter,
	r *http.Request,
	ugsn []*specialty.Ugsn,
) {
	response := make([]UgsnResponse, 0, len(ugsn))
	for _, v := range ugsn {
		response = append(response, marshalUgsnTo(*v))
	}

	render.Respond(w, r, response)
}

func marshalUgsnTo(ugsn specialty.Ugsn) UgsnResponse {
	return UgsnResponse{
		Code:  ugsn.Code(),
		Title: ugsn.Title(),
	}
}
