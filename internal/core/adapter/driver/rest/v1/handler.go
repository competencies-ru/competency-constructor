package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	logger service.Logger
	app    *app.Application
}

func NewHandler(
	app *app.Application,
	log service.Logger,
	middlewares ...rest.Middleware,
) http.Handler {
	r := chi.NewRouter()

	for _, middleware := range middlewares {
		r.Use(middleware)
	}

	return HandlerFromMux(handler{
		logger: log,
		app:    app,
	}, r)
}
