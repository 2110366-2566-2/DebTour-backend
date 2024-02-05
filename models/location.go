package models

type Location struct {
	LocationId uint    `gorm:"primaryKey;type:SERIAL" json:"location_id"`
	Name       string  `json:"name"`
	Latitude   float64 `gorm:"check:latitude >= -90 AND latitude <= 90" json:"latitude"`
	Longitude  float64 `gorm:"check:longitude >= -180 AND longitude <= 180" json:"longitude"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
}
