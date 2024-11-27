package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/arsnazarenko/basketball-service/internal/apperrors"
	"github.com/arsnazarenko/basketball-service/internal/entity"
	"github.com/arsnazarenko/basketball-service/internal/usecase"
	"github.com/arsnazarenko/basketball-service/pkg/postgres"
	"github.com/google/uuid"
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
func (p *PlayerRepo) CreatePlayer(ctx context.Context, player *entity.Player) (uuid.UUID, error) {

	query := "INSERT INTO players (name, surname, height, weight, citizenship, role, team_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	var id uuid.UUID
	if err := p.pg.Pool.QueryRow(ctx, query,
		player.Name,
		player.Surname,
		player.Height,
		player.Weight,
		player.Citizenship,
		player.Role,
		player.TeamID,
	).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("repo.CreatePlayer: create player error: %w", err)
	}
	return id, nil

}

// DeletePlayer implements usecase.PlayerRp.
func (p *PlayerRepo) DeletePlayer(ctx context.Context, playerID uuid.UUID) error {
	query := "DELETE FROM players WHERE id = $1"
	res, err := p.pg.Pool.Exec(ctx, query, playerID)
	if err != nil {
		return fmt.Errorf("repo.DeletePlayer: error: %w", err)
	}
	if res.RowsAffected() == 0 {
		return apperrors.ErrPlayerNotFound
	}
	return nil
}

// GetPlayer implements usecase.PlayerRp.
func (p *PlayerRepo) GetPlayer(ctx context.Context, playerID uuid.UUID) (*entity.Player, error) {
	query := "SELECT name, surname, height, weight, citizenship, role, team_id FROM players WHERE id = $id"

	var player entity.Player
	if err := p.pg.Pool.QueryRow(ctx, query, playerID).Scan(
		&player.Name,
		&player.Surname,
		&player.Height,
		&player.Weight,
		&player.Citizenship,
		&player.Role,
		&player.TeamID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrPlayerNotFound
		}
		return nil, fmt.Errorf("repo.GetPlayer: error: %w", err)
	}
	return &player, nil
}

// GetPlayerList implements usecase.PlayerRp.
func (p *PlayerRepo) GetPlayerList(ctx context.Context, pageSize uint64, pageNumber uint64) ([]entity.Player, error) {
	if pageNumber < 1 {
		return nil, apperrors.ErrInvalidPlayerPageNumber
	}
	if pageSize < 1 {
		return nil, apperrors.ErrInvalidPlayerPageSize
	}
	limit, offset := pageSize, (pageNumber-1)*pageSize
	query := "SELECT name, surname, height, weight, citizenship, role, team_id FROM players LIMIT $1 OFFSET $2"
	rows, err := p.pg.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrPlayerNotFound
		}
	}
	defer rows.Close()
	var list []entity.Player
	for rows.Next() {
		var player entity.Player
		if err := rows.Scan(
			&player.Name,
			&player.Surname,
			&player.Height,
			&player.Weight,
			&player.Citizenship,
			&player.Role,
			&player.TeamID,
		); err != nil {
			return nil, fmt.Errorf("repo.GetPlayerList: error: %w", err)
		}
		list = append(list, player)
	}
	return list, nil
}

// UpdatePlayer implements usecase.PlayerRp.
func (p *PlayerRepo) UpdatePlayer(ctx context.Context, playerID uuid.UUID, player *entity.Player) error {
	query := "UPDATE players SET name = $1, surname = $2, height = $3, weight = $4, citizenship = $5, role = $6, team_id = $7 WHERE id = $8"

	res, err := p.pg.Pool.Exec(ctx, query,
		player.Name,
		player.Surname,
		player.Height,
		player.Weight,
		player.Citizenship,
		player.Role,
		player.TeamID,
		playerID,
	)
	if err != nil {
		return fmt.Errorf("repo.UpdatePlayer: error: %w", err)
	}
	if res.RowsAffected() == 0 {
		return apperrors.ErrPlayerNotFound
	}
	return nil
}
