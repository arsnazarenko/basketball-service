package app

import (
	"context"

	"github.com/arsnazarenko/basketball-service/config"
	"github.com/arsnazarenko/basketball-service/pkg/logger"
	"github.com/arsnazarenko/basketball-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Postgresql
	psql, err := postgres.New(cfg.Postgres.PostgresURL)
	if err != nil {
		panic(err)
	}
	l.Info("PostgreSQL is started and connected by URL: %s", cfg.Postgres.PostgresURL)

	var greeting string
	err = psql.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		l.Error("QueryRow failed: %v\n", err)
	}
	l.Info(greeting)
	psql.Close()

}
