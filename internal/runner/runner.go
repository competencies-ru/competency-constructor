package runner

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/competencies-ru/competency-constructor/pkg/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/competencies-ru/competency-constructor/internal/config"
	"github.com/competencies-ru/competency-constructor/internal/server"
)

type singletonPostgres struct {
	db  *pgxpool.Pool
	one sync.Once
}

type Runner struct {
	singletonPostgres

	cfg    *config.Config
	server *server.Server
}

func New(path string) *Runner {
	r := &Runner{}

	r.initConfig(path)
	r.initServer()
	r.initPersistent()

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

func (r *Runner) initPersistent() {
	_ = r.postgres()

	log.Println("postgres init")
}

func (r *Runner) initServer() {
	log.Println("init server")

	r.server = server.NewServer(r.cfg.HTTP, nil)
}

func (r *Runner) postgres() *pgxpool.Pool {
	r.singletonPostgres.one.Do(func() {
		client, err := postgres.NewClient(r.cfg.Postgres)
		if err != nil {
			log.Println("error connection database", err)
		}

		r.singletonPostgres.db = client
	})

	return r.singletonPostgres.db
}

func (r *Runner) StartServer() {
	if err := r.server.Start(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("HTTP Server stopped with an error\n", err)
	}
}

func (r *Runner) Stop() {
	r.shutdownServer()
	r.disconnectPostgres()
}

func (r *Runner) shutdownServer() {
	log.Println("Start shutdown server...")

	ctx, stop := context.WithTimeout(context.Background(), r.cfg.HTTP.ShutdownTimeout)
	defer stop()

	if err := r.server.Shutdown(ctx); err != nil {
		log.Println("error shutdown server", err)
	}
}

func (r *Runner) disconnectPostgres() {
	r.singletonPostgres.db.Close()
}
