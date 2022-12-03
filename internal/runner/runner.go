package runner

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	v1 "github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest/v1"
	"github.com/competencies-ru/competency-constructor/internal/core/app"

	zapAdapter "github.com/competencies-ru/competency-constructor/internal/core/adapter/driven/logger/zap"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/competencies-ru/competency-constructor/pkg/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/competencies-ru/competency-constructor/internal/config"
	"github.com/competencies-ru/competency-constructor/internal/server"
)

type singletonPostgres struct {
	db  *pgxpool.Pool
	one sync.Once
}

type singletonZapLogger struct {
	one    sync.Once
	logger *zap.Logger
}

type Runner struct {
	singletonZapLogger
	singletonPostgres

	logger service.Logger
	config *config.Config
	server *server.Server
	app    *app.Application
}

func New(path string) *Runner {
	r := &Runner{}

	r.initConfig(path)
	r.initLogger()
	r.initServer()
	r.initPersistent()
	r.initApplication()

	return r
}

func (r *Runner) initConfig(path string) {
	log.Println("init config to path: " + path)

	cfg, err := config.ParseFrom(path)
	if err != nil {
		log.Fatalf("config parsing error: %v", err)
	}

	r.config = cfg
}

func (r *Runner) initPersistent() {
	_ = r.postgres()
}

func (r *Runner) initLogger() {
	if r.config.Logger.Lib == config.Zap {
		r.logger = zapAdapter.NewLogger(r.zap())
	}

	r.logger.Info(
		"Logger library used",
		"lib", r.config.Logger.Lib,
		"level", r.config.Logger.Level,
	)
}

func (r *Runner) initServer() {
	r.logger.Info("start init server")

	r.server = server.NewServer(r.config.HTTP, rest.NewHandler(rest.Params{
		Middlewares: rest.Middlewares(r.config.HTTP.AllowedOrigins),
		Routes: []rest.Route{
			{
				Pattern: "/competency-constructor/api/v1",
				Handler: v1.NewHandler(r.app, r.logger),
			},
		},
	}))
}

func (r *Runner) initApplication() {
	r.app = &app.Application{}
}

func (r *Runner) postgres() *pgxpool.Pool {
	r.singletonPostgres.one.Do(func() {
		client, err := postgres.NewClient(r.config.Postgres)
		if err != nil {
			r.logger.Fatal("failed to connect to database", err)
		}

		r.singletonPostgres.db = client
	})

	return r.singletonPostgres.db
}

func (r *Runner) zap() *zap.Logger {
	r.singletonZapLogger.one.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(r.levelLogger())

		build, err := cfg.Build()
		if err != nil {
			log.Fatalf("failed to configure logger: %v", err)
		}

		r.singletonZapLogger.logger = build
	})

	return r.singletonZapLogger.logger
}

func (r *Runner) StartServer() {
	if err := r.server.Start(); !errors.Is(err, http.ErrServerClosed) {
		r.logger.Fatal("HTTP Server stopped with an error", err)
	}
}

func (r *Runner) Stop() {
	r.shutdownServer()
	r.disconnectPostgres()
}

func (r *Runner) shutdownServer() {
	r.logger.Info("Start shutdown server...")

	ctx, stop := context.WithTimeout(context.Background(), r.config.HTTP.ShutdownTimeout)
	defer stop()

	if err := r.server.Shutdown(ctx); err != nil {
		r.logger.Fatal("error shutdown server", err)
	}
}

func (r *Runner) disconnectPostgres() {
	r.singletonPostgres.db.Close()
}

func (r *Runner) levelLogger() zapcore.Level {
	switch r.config.Logger.Level {
	case config.DebugLevel:
		return zapcore.DebugLevel
	case config.InfoLevel:
		return zapcore.InfoLevel
	case config.WarnLevel:
		return zapcore.WarnLevel
	case config.ErrorLevel:
		return zapcore.ErrorLevel
	case config.FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
