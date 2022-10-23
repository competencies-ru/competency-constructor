package server

import (
	"context"
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/config"
)

type Server struct {
	server http.Server
}

func NewServer(cfg config.HTTP, handler http.Handler) *Server {
	return &Server{
		server: http.Server{
			Addr:         ":" + cfg.Port,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			Handler:      handler,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
