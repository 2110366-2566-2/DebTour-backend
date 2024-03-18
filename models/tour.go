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
	AgencyUsername   string    `gorm:"not null" json:"agencyUsername"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime" json:"createdTimestamp"`
}

type TourWithActivitiesWithLocationWithImagesRequest struct {
	Name             string                        `json:"name"`
	StartDate        string                        `json:"startDate"`
	EndDate          string                        `json:"endDate"`
	Description      string                        `json:"description"`
	OverviewLocation string                        `json:"overviewLocation"`
	Price            float64                       `json:"price"`
	RefundDueDate    string                        `json:"refundDueDate"`
	MaxMemberCount   uint                          `json:"maxMemberCount"`
	Activities       []ActivityWithLocationRequest `json:"activities"`
	Images			 []string						`json:"images"`
}

func ToTour(tourRequest TourWithActivitiesWithLocationWithImagesRequest, tourId uint, agencyUsername string) (Tour, error) {
	return Tour{
		TourId:           tourId,
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
		AgencyUsername:   agencyUsername,
	}, nil
}

type TourWithActivitiesWithLocationWithImages struct {
	TourId           uint                   `gorm:"primaryKey;autoIncrement" json:"tourId"`
	Name             string                 `gorm:"not null" json:"name"`
	StartDate        string                 `gorm:"not null" json:"startDate"`
	EndDate          string                 `gorm:"not null" json:"endDate"`
	Description      string                 `gorm:"not null" json:"description"`
	OverviewLocation string                 `gorm:"not null" json:"overviewLocation"`
	Price            float64                `gorm:"not null;check:price > 0" json:"price"`
	RefundDueDate    string                 `gorm:"not null" json:"refundDueDate"`
	MaxMemberCount   uint                   `gorm:"not null" json:"maxMemberCount"`
	MemberCount      uint                   `gorm:"not null" json:"memberCount"`
	Status           string                 `gorm:"not null" json:"status"`
	AgencyUsername   string                 `gorm:"foreignKey:UserRefer" json:"agencyUsername"`
	Activities       []ActivityWithLocation `json:"activities"`
	Images			 []string				 `json:"images"`
}

func ToTourWithActivitiesWithLocationWithImages(tour Tour, activities []ActivityWithLocation, images []string) (TourWithActivitiesWithLocationWithImages, error) {
	return TourWithActivitiesWithLocationWithImages{
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
		Activities:       activities,
		Images:			  images,
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
