package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Params struct {
		Middlewares []Middleware
		Routes      []Route
	}

	Middleware func(next http.Handler) http.Handler

	Route struct {
		Pattern string
		Handler http.Handler
	}
)

func NewHandler(p Params) http.Handler {
	root := chi.NewRouter()

	for _, middleware := range p.Middlewares {
		root.Use(middleware)
	}

	for _, route := range p.Routes {
		root.Mount(route.Pattern, route.Handler)
	}

	return root
}
