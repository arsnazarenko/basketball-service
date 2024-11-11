package entity

import "time"

type Game struct {
	HomeTeamID  string    `json:"homeTeamId,omitempty" default:""`
	GuestTeamID string    `json:"guestTeamId,omitempty" default:""`
	Date        time.Time `json:"date,omitempty" default:""`
	GameType    string    `json:"gameType,omitempty" default:""`
	Winner      uint8     `json:"winner,omitempty" default:""`
}
