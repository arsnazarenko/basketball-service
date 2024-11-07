package entity

type Statistics struct {
	PlayerID      string `json:"playerId,omitempty"`
	GameID        string `json:"gameId,omitempty"`
	Points        int    `json:"points,omitempty"`
	Assists       int    `json:"assists,omitempty"`
	Rebounds      int    `json:"rebounds,omitempty"`
	Interceptions int    `json:"interceptions,omitempty"`
}
