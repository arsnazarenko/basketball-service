package entity

type Conference string

const (
	ConferenceEastern = "EASTERN"
	ConferenceWester  = "WESTERN"
)

type Team struct {
	Name       string     `json:"name,omitempty"`
	Conference Conference `json:"conference,omitempty"`
}
