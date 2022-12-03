package rest

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const maxAge = 300

func Middlewares(allowedOriginsCors []string) []Middleware {
	return []Middleware{
		middleware.RealIP,
		middleware.Recoverer,
		middleware.NoCache,
		middleware.DefaultLogger,
		corsMiddleware(allowedOriginsCors),
	}
}

func corsMiddleware(allowedOrigins []string) Middleware {
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}).Handler
}
