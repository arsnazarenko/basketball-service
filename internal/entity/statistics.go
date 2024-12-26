package entity

import "github.com/google/uuid"


type Statistics struct {
	PlayerID      uuid.UUID `json:"playerId,omitempty"`
	GameID        uuid.UUID `json:"gameId,omitempty"`
	Points        int    `json:"points,omitempty"`
	Assists       int    `json:"assists,omitempty"`
	Rebounds      int    `json:"rebounds,omitempty"`
	Interceptions int    `json:"interceptions,omitempty"`
}
