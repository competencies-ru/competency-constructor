package v1

import (
	"net/http"

	"github.com/google/uuid"

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

func toUUIDPointer(s string) *uuid.UUID {
	if s == "" {
		return nil
	}

	parse := uuid.MustParse(s)

	return &parse
}

func toUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}

func toString(uuid *uuid.UUID) string {
	if uuid == nil {
		return ""
	}

	return uuid.String()
}
