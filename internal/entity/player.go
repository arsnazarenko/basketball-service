package entity

type Player struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Age         uint8  `json:"age,omitempty"`
	Height      uint8  `json:"height,omitempty"`
	Weight      uint8  `json:"weight,omitempty"`
	Citizenship string `json:"citizenship,omitempty"`
	Role        string `json:"role,omitempty"`
	TeamID      string `json:"teamId,omitempty"`
}
