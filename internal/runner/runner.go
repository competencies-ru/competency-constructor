package runner

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/competencies-ru/competency-constructor/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/multierr"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	v1 "github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest/v1"
	"github.com/competencies-ru/competency-constructor/internal/core/app"

	zapAdapter "github.com/competencies-ru/competency-constructor/internal/core/adapter/driven/logger/zap"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/competencies-ru/competency-constructor/internal/config"
	"github.com/competencies-ru/competency-constructor/internal/server"
)

type singletonMongodb struct {
	one sync.Once
	db  *mongo.Database
}

type singletonZapLogger struct {
	one    sync.Once
	logger *zap.Logger
}

type persistenceContext struct {
	ugsnRepo service.UgsnRepository
}

type Runner struct {
	singletonZapLogger
	singletonMongodb

	logger  service.Logger
	config  *config.Config
	server  *server.Server
	app     *app.Application
	pcontex persistenceContext
}

func New(path string) *Runner {
	r := &Runner{}

	r.initConfig(path)
	r.initLogger()
	r.initPersistent()
	r.initApplication()
	r.initServer()

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
	_ = r.mongo()
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
	r.app = &app.Application{
		Services: app.Services{
			UgsnService: service.NewUgsnHandler(r.pcontex.ugsnRepo),
		},
	}
}

func (r *Runner) mongo() *mongo.Database {
	r.singletonMongodb.one.Do(func() {
		client, err := mongodb.NewClient(
			r.config.Mongodb.URI,
			r.config.Mongodb.Username,
			r.config.Mongodb.Password,
		)
		if err != nil {
			r.logger.Fatal("failed to connect to database", err)
		}

		r.singletonMongodb.db = client.Database(r.config.Mongodb.DatabaseName)
	})

	return r.singletonMongodb.db
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
	err := multierr.Append(
		r.shutdownServer(),
		r.disconnectMongo(),
	)
	if err != nil {
		r.logger.Fatal("Runner stopped with an error", err)
	}
}

func (r *Runner) shutdownServer() error {
	r.logger.Info("Start shutdown server...")

	ctx, stop := context.WithTimeout(context.Background(), r.config.HTTP.ShutdownTimeout)
	defer stop()

	return r.server.Shutdown(ctx)
}

func (r *Runner) disconnectMongo() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		r.config.Mongodb.DisconnectTimeout,
	)
	defer cancel()

	return r.mongo().Client().Disconnect(ctx)
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
