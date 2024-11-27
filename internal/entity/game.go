package entity

import "time"

type Game struct {
	HomeTeamID  string    `json:"homeTeamId,omitempty" default:""`
	GuestTeamID string    `json:"guestTeamId,omitempty" default:""`
	GameDate    time.Time `json:"gameDate,omitempty" default:""`
	HomeScore   uint8     `json:"homeScore,omitempty" default:""`
	GuestScore   uint8     `json:"guestScore,omitempty" default:""`
}
