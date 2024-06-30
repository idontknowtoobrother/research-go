package models

type Player struct {
	Uuid       string  `json:"uuid"`
	Name       string  `json:"name"`
	Experience float64 `json:"experience"`
	Inventory  []Item  `json:"inventory"`
}
