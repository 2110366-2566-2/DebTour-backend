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
	LocationId uint    `json:"locationId"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
}

func ToLocation(locationRequest LocationRequest) Location {
	return Location{
		Name:      locationRequest.Name,
		Latitude:  locationRequest.Latitude,
		Longitude: locationRequest.Longitude,
		Type:      locationRequest.Type,
		Address:   locationRequest.Address,
	}
}

func GetAllLocations() (locations []Location, err error) {
	result := db.Model(&Location{}).Find(&locations)

	return locations, result.Error
}

func GetLocationById(locationId uint) (Location, error) {
	var location Location
	result := db.Model(&Location{}).First(&location, locationId)
	return location, result.Error
}

func CreateLocation(location *Location) (err error) {
	result := db.Model(&Location{}).Create(location)

	return result.Error
}

func UpdateLocation(location *Location) (err error) {
	_, err = GetLocationById(location.LocationId)

	if err != nil {
		return err
	}

	result := db.Model(&Location{}).Where("location_id = ?", location.LocationId).Updates(location)

	return result.Error
}

func DeleteLocation(locationId uint) (err error) {
	result := db.Model(&Location{}).Where("location_id = ?", locationId).Delete(&Location{})

	return result.Error
}
