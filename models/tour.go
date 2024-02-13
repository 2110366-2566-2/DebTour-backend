package models

import (
	"time"
)

type Tour struct {
	TourId           uint      `gorm:"primaryKey;autoIncrement" json:"tourId"`
	Name             string    `gorm:"not null" json:"name"`
	StartDate        string    `gorm:"not null" json:"startDate"`
	EndDate          string    `gorm:"not null" json:"endDate"`
	Description      string    `gorm:"not null" json:"description"`
	OverviewLocation string    `gorm:"not null" json:"overviewLocation"`
	Price            float64   `gorm:"not null;check:price > 0" json:"price"`
	RefundDueDate    string    `gorm:"not null" json:"refundDueDate"`
	MaxMemberCount   uint      `gorm:"not null" json:"maxMemberCount"`
	MemberCount      uint      `gorm:"not null" json:"memberCount"`
	Status           string    `gorm:"not null" json:"status"`
	AgencyUsername   string    `gorm:"foreignKey:UserRefer" json:"agencyUsername"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime" json:"createdTimestamp"`
}

type TourRequest struct {
	Name             string            `json:"name"`
	StartDate        string            `json:"startDate"`
	EndDate          string            `json:"endDate"`
	Description      string            `json:"description"`
	OverviewLocation string            `json:"overviewLocation"`
	Price            float64           `json:"price"`
	RefundDueDate    string            `json:"refundDueDate"`
	MaxMemberCount   uint              `json:"maxMemberCount"`
	Activities       []ActivityRequest `json:"activities"`
}

type TourActivityLocation struct {
	TourId           uint               `json:"tourId"`
	Name             string             `json:"name"`
	StartDate        string             `json:"startDate"`
	EndDate          string             `json:"endDate"`
	Description      string             `json:"description"`
	OverviewLocation string             `json:"overviewLocation"`
	Price            float64            `json:"price"`
	RefundDueDate    string             `json:"refundDueDate"`
	MaxMemberCount   uint               `json:"maxMemberCount"`
	MemberCount      uint               `json:"memberCount"`
	Status           string             `json:"status"`
	AgencyUsername   string             `json:"agencyUsername"`
	CreatedTimestamp time.Time          `json:"createdTimestamp"`
	Activities       []ActivityResponse `json:"activities"`
}

func ToTour(tourRequest TourRequest) Tour {
	return Tour{
		Name:             tourRequest.Name,
		StartDate:        tourRequest.StartDate,
		EndDate:          tourRequest.EndDate,
		Description:      tourRequest.Description,
		OverviewLocation: tourRequest.OverviewLocation,
		Price:            tourRequest.Price,
		RefundDueDate:    tourRequest.RefundDueDate,
		MaxMemberCount:   tourRequest.MaxMemberCount,
		MemberCount:      0,
		Status:           "Available",
	}
}

func ToTourActivityLocation(tour Tour, activities []Activity) (TourActivityLocation, error) {
	var activityResponses []ActivityResponse
	for _, activity := range activities {
		activityResponse, err := ToActivityResponse(activity)
		if err != nil {
			return TourActivityLocation{}, err
		}
		activityResponses = append(activityResponses, activityResponse)
	}

	return TourActivityLocation{
		TourId:           tour.TourId,
		Name:             tour.Name,
		StartDate:        tour.StartDate,
		EndDate:          tour.EndDate,
		Description:      tour.Description,
		OverviewLocation: tour.OverviewLocation,
		Price:            tour.Price,
		RefundDueDate:    tour.RefundDueDate,
		MaxMemberCount:   tour.MaxMemberCount,
		MemberCount:      tour.MemberCount,
		Status:           tour.Status,
		AgencyUsername:   tour.AgencyUsername,
		CreatedTimestamp: tour.CreatedTimestamp,
		Activities:       activityResponses,
	}, nil
}

type FilteredToursResponse struct {
	TourId           int     `json:"tourId"`
	TourName         string  `json:"tourName"`
	StartDate        string  `json:"startDate"`
	EndDate          string  `json:"endDate"`
	OverviewLocation string  `json:"overviewLocation"`
	MemberCount      uint    `json:"memberCount"`
	MaxMemberCount   uint    `json:"maxMemberCount"`
	Price            float64 `json:"price"`
}

func GetAllTours() (tours []Tour, err error) {
	result := db.Find(&tours)

	return tours, result.Error
}

func GetTourById(tourId int) (TourActivityLocation, error) {
	var tour Tour
	result := db.First(&tour, tourId)

	var tourActivityLocation TourActivityLocation

	if result.Error != nil {
		return tourActivityLocation, result.Error
	}

	acitivities, err := GetAllActivitiesByTourId(tour.TourId)

	if err != nil {
		return tourActivityLocation, err
	}

	tourActivityLocation, err = ToTourActivityLocation(tour, acitivities)

	return tourActivityLocation, err
}

func GetOnlyTourById(tourId int) (Tour, error) {
	var tour Tour
	result := db.First(&tour, tourId)

	return tour, result.Error
}

func CreateTour(tour *Tour, activitiesRequest []ActivityRequest) (err error) {
	tx := db.Begin()

	result := db.Model(&Tour{}).Create(tour)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	for _, activityRequest := range activitiesRequest {
		activity := ToActivity(activityRequest, tour.TourId)
		err = CreateActivity(&activity, activityRequest.Location, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func UpdateTour(tour *Tour) (err error) {
	_, err = GetTourById(int(tour.TourId))

	if err != nil {
		return err
	}

	result := db.Model(&Tour{}).Where("tour_id = ?", tour.TourId).Updates(tour)

	return result.Error
}

func DeleteTour(tourId uint) (err error) {
	// Transaction begins
	tx := db.Begin()

	// Delete all joinings of the tour by calling the function from joining.go
	err = DeleteJoiningByTourId(tourId, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete the tour
	result := tx.Model(&Tour{}).Where("tour_id = ?", tourId).Delete(&Tour{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Transaction ends
	tx.Commit()
	return nil
}

func FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo string, offset, limit int) ([]Tour, error) {
	var tours []Tour
	result := db.Model(&Tour{}).Select("tour_id, name, start_date, end_date, overview_location, member_count, max_member_count, price").Where("name LIKE ? AND start_date >= ? AND end_date <= ? AND overview_location LIKE ? AND member_count >= ? AND member_count <= ? AND price >= ? AND price <= ?", name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo).Limit(limit).Offset(offset).Find(&tours)
	return tours, result.Error
}
