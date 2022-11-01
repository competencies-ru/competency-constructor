package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/competencies-ru/competency-constructor/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
)

// NewClient creates a new pool connection to the postgres database.
func NewClient(cfg config.Postgres) (*pgxpool.Pool, error) {
	ctx := context.Background()
	url := initURLPostgres(cfg)

	parseConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	parseConfig.MaxConns = maxConn
	parseConfig.HealthCheckPeriod = healthCheckPeriod
	parseConfig.MaxConnIdleTime = maxConnIdleTime
	parseConfig.MaxConnLifetime = maxConnLifetime
	parseConfig.MinConns = minConns

	pool, err := pgxpool.NewWithConfig(ctx, parseConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func initURLPostgres(cfg config.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.UserName,
		cfg.DataBaseName,
		cfg.Password,
	)
}
