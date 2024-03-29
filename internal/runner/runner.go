package runner

import (
	"context"
	"errors"
	"github.com/competencies-ru/competency-constructor/pkg/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"sync"

	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"

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
	postgresRepo "github.com/competencies-ru/competency-constructor/internal/core/adapter/driven/persistence/postgres"
	"github.com/competencies-ru/competency-constructor/internal/server"
)

type singletonMongodb struct {
	one sync.Once
	db  *mongo.Database
}

type singletonPostgres struct {
	one sync.Once
	db  *pgxpool.Pool
}

type singletonZapLogger struct {
	one    sync.Once
	logger *zap.Logger
}

type persistenceContext struct {
	levelRepo            service.LevelRepository
	levelsReadModel      query.LevelReadModels
	ugsnReadModels       query.UgsnReadModels
	specialtyReadModels  query.SpecialtiesReadModels
	programReadModels    query.ProgramsReadModels
	ugsnRepo             service.UgsnRepository
	specialtyRepo        service.SpecialtyRepository
	programRepo          service.ProgramRepository
	competencyRepo       service.CompetencyRepository
	filterCompetencyRepo query.FilterCompetencyReadModels
}

type Runner struct {
	singletonZapLogger
	singletonMongodb
	singletonPostgres

	logger      service.Logger
	config      *config.Config
	server      *server.Server
	app         *app.Application
	persistence persistenceContext
}

func New(path string) *Runner {
	r := &Runner{}

	r.initConfig(path)
	r.initLogger()
	r.initPostgresPersistent()
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

func (r *Runner) initPostgresPersistent() {
	database := r.postgres()
	args := []interface{}{"db", "postgres"}

	levelRepo := postgresRepo.NewLevelRepository(database)
	ugsnRepo := postgresRepo.NewUgsnRepository(database)
	specialtyRepo := postgresRepo.NewSpecialtyRepository(database)
	programRepo := postgresRepo.NewProgramRepository(database)
	competencyRepo := postgresRepo.NewCompetencyRepository(database)

	r.persistence.levelRepo = levelRepo
	r.logger.Info("Level repository initialization completed", args...)

	r.persistence.levelsReadModel = levelRepo
	r.logger.Info("Levels read model repository initialization completed", args...)

	r.persistence.ugsnRepo = ugsnRepo
	r.logger.Info("Ugsn repository initialization completed", args...)

	r.persistence.specialtyRepo = specialtyRepo
	r.logger.Info("Specialty repository initialization completed", args...)

	r.persistence.programRepo = programRepo
	r.logger.Info("Program repository initialization completed", args...)

	r.persistence.programReadModels = programRepo
	r.logger.Info("Program read models repository initialization completed", args...)

	r.persistence.specialtyReadModels = specialtyRepo
	r.logger.Info("Specialty read models repository initialization completed", args...)

	r.persistence.ugsnReadModels = ugsnRepo
	r.logger.Info("Ugsn read models repository initialization completed", args...)

	r.persistence.competencyRepo = competencyRepo
	r.logger.Info("Competency repository initialization completed", args...)

	r.persistence.filterCompetencyRepo = competencyRepo
	r.logger.Info("filter competency models repository initialization completed", args...)

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
			{
				Pattern: "/swagger-ui",
				Handler: http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger/v1"))),
			},
		},
	}))
}

func (r *Runner) initApplication() {
	r.app = &app.Application{
		Commands: app.Commands{
			CreateLevel:    command.NewCreateLevelHandler(r.persistence.levelRepo),
			AddUgsn:        command.NewAddUgsnHandler(r.persistence.ugsnRepo, r.persistence.levelRepo),
			AddPrograms:    command.NewAddProgramsHandler(r.persistence.programRepo, r.persistence.specialtyRepo),
			AddSpecialties: command.NewAddSpecialtiesHandler(r.persistence.specialtyRepo, r.persistence.ugsnRepo),
			CreateCompetency: command.NewCreateCompetencyHandler(
				r.persistence.competencyRepo,
				r.persistence.levelRepo,
				r.persistence.ugsnRepo,
				r.persistence.specialtyRepo,
				r.persistence.programRepo,
			),
		},
		Queries: app.Queries{
			FindLevels:         query.NewFindLevelsHandler(r.persistence.levelsReadModel),
			FindAllUgsn:        query.NewFindUgsnHandler(r.persistence.ugsnReadModels),
			FindAllSpecialties: query.NewFindSpecialtiesHandler(r.persistence.specialtyReadModels),
			FindAllPrograms:    query.NewFindProgramsHandler(r.persistence.programReadModels),
			FindAllCompetency:  query.NewFilterCompetenciesHandler(r.persistence.filterCompetencyRepo),
		},
	}
}

func (r *Runner) postgres() *pgxpool.Pool {
	r.singletonPostgres.one.Do(func() {
		client, err := postgres.NewClient(r.config.Postgres)
		if err != nil {
			log.Fatalf("failed to connect to postgres database", err)
		}

		r.singletonPostgres.db = client
	})

	return r.singletonPostgres.db
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
