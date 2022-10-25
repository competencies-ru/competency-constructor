package runner

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/config"
	"github.com/competencies-ru/competency-constructor/internal/server"
)

type Runner struct {
	cfg    *config.Config
	server *server.Server
}

func New(path string) *Runner {
	r := &Runner{}

	r.initConfig(path)
	r.initServer()

	return r
}

func (r *Runner) initConfig(path string) {
	log.Println("init config to path: " + path)

	cfg, err := config.Parse(path)
	if err != nil {
		log.Fatal("config parsing error", err)
	}

	r.cfg = cfg
}

func (r *Runner) initServer() {
	log.Println("init server")

	r.server = server.NewServer(r.cfg.HTTP, nil)
}

func (r *Runner) StartServer() {
	if err := r.server.Start(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("HTTP Server stopped with an error", err)
	}
}

func (r *Runner) Stop() {
	log.Println("Start shutdown server...")

	ctx, stop := context.WithTimeout(context.Background(), r.cfg.HTTP.ShutdownTimeout)
	defer stop()

	if err := r.server.Shutdown(ctx); err != nil {
		log.Println("error shutdown server", err)
	}
}
