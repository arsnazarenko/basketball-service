package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/arsnazarenko/basketball-service/internal/apperrors"
	"github.com/arsnazarenko/basketball-service/internal/entity"
	"github.com/arsnazarenko/basketball-service/internal/usecase"
	"github.com/arsnazarenko/basketball-service/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

var _ usecase.PlayerRp = (*PlayerRepo)(nil)

type PlayerRepo struct {
	pg *postgres.Postgres
}

func NewPlayerRepo(pg *postgres.Postgres) *PlayerRepo {
	return &PlayerRepo{
		pg: pg,
	}
}

// CreatePlayer implements usecase.PlayerRp.
func (p *PlayerRepo) CreatePlayer(ctx context.Context, player *entity.Player) (string, error) {

	query := "INSERT INTO players (name, surname, age, height, weight, citizenship, role, team_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	var id string
	if err := p.pg.Pool.QueryRow(ctx, query,
		player.Name,
		player.Surname,
		player.Age,
		player.Height,
		player.Weight,
		player.Citizenship,
		player.Role,
		player.TeamID,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("repo.CreatePlayer: create player error: %w", err)
	}
	return id, nil

}

// DeletePlayer implements usecase.PlayerRp.
func (p *PlayerRepo) DeletePlayer(ctx context.Context, playerID string) error {
	query := "DELETE FROM players WHEN id = $1 RETURNING id"
	var id string
	if err := p.pg.Pool.QueryRow(ctx, query, playerID).Scan(&id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return apperrors.ErrInvalidPlayerID
		}
		return fmt.Errorf("repo.DeletePlayer: error: %w", err)

	}
	return nil
}

// GetPlayer implements usecase.PlayerRp.
func (p *PlayerRepo) GetPlayer(ctx context.Context, playerID string) (*entity.Player, error) {
	query := "SELECT * FROM players WHERE id = $id"

	var (
		player entity.Player
		id string
	)
	if err := p.pg.Pool.QueryRow(ctx, query, playerID).Scan(&id, &player.Name, &player.Surname, &player.Age, &player.Height, &player.Weight, &player.Citizenship, &player.Role, &player.TeamID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrPlayerNotFound
		}
		return nil, fmt.Errorf("repo.GetPlayer: error: %w", err)
	}
	return &player, nil
}

// GetPlayerList implements usecase.PlayerRp.
func (p *PlayerRepo) GetPlayerList(ctx context.Context, pageSize int64, pageNumber int64) ([]*entity.Player, error) {
	panic("unimplemented")
}

// UpdatePlayer implements usecase.PlayerRp.
func (p *PlayerRepo) UpdatePlayer(ctx context.Context, playerID string, player *entity.Player) (*entity.Player, error) {
	panic("unimplemented")
}
