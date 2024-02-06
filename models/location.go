package models

type Location struct {
	LocationId uint    `gorm:"primaryKey;autoIncrement" json:"locationId"`
	Name       string  `json:"name"`
	Latitude   float64 `gorm:"check:latitude >= -90 AND latitude <= 90" json:"latitude"`
	Longitude  float64 `gorm:"check:longitude >= -180 AND longitude <= 180" json:"longitude"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
}
