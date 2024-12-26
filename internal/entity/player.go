package entity

import "github.com/jackc/pgx/v5/pgtype"

type Role string

const (
	RolePointGuard    = "PG"
	RoleShootingGuard = "SG"
	RoleSmallForward  = "SF"
	RolePowerForward  = "PF"
	RoleCenter        = "PF"
)

type Player struct {
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"surname,omitempty"`
	Height      uint8     `json:"height,omitempty"`
	Weight      uint8     `json:"weight,omitempty"`
	Citizenship string    `json:"citizenship,omitempty"`
	Role        Role      `json:"role,omitempty"`
	TeamID      pgtype.UUID `json:"teamId,omitempty"`
}
