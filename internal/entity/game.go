package entity

type Game struct {
	HomeTeamID  string `json:"homeTeamId,omitempty" default:""`
	GuestTeamID string `json:"guestTeamId,omitempty" default:""`
	Date        string `json:"date,omitempty" default:""`
	Type        string `json:"type,omitempty" default:""`
	Winner      uint8  `json:"winner,omitempty" default:""`
}
