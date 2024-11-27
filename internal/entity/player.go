package entity

import uuid "github.com/vgarvardt/pgx-google-uuid/v5"

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
	TeamID      uuid.UUID `json:"teamId,omitempty"`
}
