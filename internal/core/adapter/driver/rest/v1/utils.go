package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/go-chi/render"
)

func decode[T any](w http.ResponseWriter, r *http.Request, v T) bool {
	if err := render.Decode(r, v); err != nil {
		rest.BadRequest(string(BadRequest), err, w, r)

		return false
	}

	return true
}
