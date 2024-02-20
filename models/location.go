package models

type Location struct {
	LocationId uint    `gorm:"primaryKey;autoIncrement" json:"locationId"`
	Name       string  `gorm:"not null" json:"name"`
	Latitude   float64 `gorm:"not null;check:latitude >= -90 AND latitude <= 90" json:"latitude"`
	Longitude  float64 `gorm:"not null;check:longitude >= -180 AND longitude <= 180" json:"longitude"`
	Type       string  `gorm:"not null" json:"type"`
	Address    string  `gorm:"not null" json:"address"`
}

type LocationRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
	Address   string  `json:"address"`
}

func ToLocation(locationRequest LocationRequest, locationId uint) Location {
	return Location{
		LocationId: locationId,
		Name:      locationRequest.Name,
		Latitude:  locationRequest.Latitude,
		Longitude: locationRequest.Longitude,
		Type:      locationRequest.Type,
		Address:   locationRequest.Address,
	}
}