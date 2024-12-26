package usecase

import (
	"context"

	"github.com/arsnazarenko/basketball-service/internal/entity"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	// Player - use case
	Player interface {
		CreatePlayer(ctx context.Context, player *entity.Player) (string, error)
		UpdatePlayer(ctx context.Context, playerID int, player *entity.Player) (*entity.Player, error)
		GetPlayer(ctx context.Context, playerID int) (*entity.Player, error)
		DeletePlayer(ctx context.Context, playerID int) error
		GetPlayerList(ctx context.Context) ([]*entity.Player, error)
	}

	// PlayerRp - mongodb
	PlayerRp interface {
		CreatePlayer(ctx context.Context, player *entity.Player) (pgtype.UUID, error)
		UpdatePlayer(ctx context.Context, playerId pgtype.UUID, player *entity.Player) error
		DeletePlayer(ctx context.Context, playerId pgtype.UUID) error
		GetPlayer(ctx context.Context, playerId pgtype.UUID) (*entity.Player, error)
		GetPlayerList(ctx context.Context, count, offset uint64) ([]entity.Player, error)
	}
)
