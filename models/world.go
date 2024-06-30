package models

type World struct {
	Uuid    string   `json:"uuid"`
	Players []Player `json:"players"`
}
