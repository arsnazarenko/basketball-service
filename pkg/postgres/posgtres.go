package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres.NewPostgres.pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)
	poolConfig.AfterConnect = createTables
	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres.NewPostgres.connAttempts == 0: %w", err)
	}

	return pg, nil
}
func createTables(ctx context.Context, conn *pgx.Conn) error {
	const tableCreationQuery = `
CREATE TYPE role AS ENUM ('PG', 'SG', 'SF', 'PF', 'C');
CREATE TYPE conference AS ENUM ('EASTERN', 'WESTERN');

CREATE TABLE IF NOT EXISTS players (
    id uuid not null default gen_random_uuid(),
    name varchar(256) not null,
    surname varchar(256) not null,
    height smallint not null CHECK (height > 100 AND height < 310),
    weight smallint not null CHECK (weight > 20 AND weight < 400),
    citizenship varchar(256) not null,
    role role not null,
    team_id uuid not null references teams(id),
    primary_key(id)
)


CREATE TABLE IF NOT EXISTS teams (
    id uuid not null default gen_random_uuid(),
    name varchar(256) not null,
    conf conference not null,
    primary_key(id)
)


CREATE TABLE IF NOT EXISTS games (
    id uuid not null default gen_random_uuid(),
    home_team_id uuid not null references teams(id),
    guest_team_id uuid not null references teams(id),
    game_date date not null,
    home_score smallint not null,
    guest_score smallint not null,
    primary_key(id)
)

CREATE TABLE IF NOT EXISTS statistics (
    player_id uuid not null references players(id),
    game_id uuid not null references games(id),
    points smallint not null,
    assists smallint not null,
    rebound smallint not null,
    interceptions smallint not null,
    primary_key(player_id, game_id)
)`

	if _, err := conn.Exec(ctx, tableCreationQuery); err != nil {
		return fmt.Errorf("postgres.createTables: %w", err)
	}
	return nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
