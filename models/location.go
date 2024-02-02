package models

type Location struct {
	LocationId uint    `json:"locationId"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
}
