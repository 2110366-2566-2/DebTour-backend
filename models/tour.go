package models

import "time"

type Tour struct {
	TourId           uint      `gorm:"primaryKey;autoIncrement" json:"tourId"`
	Name             string    `json:"name"`
	StartDate        string    `json:"startDate"`
	EndDate          string    `json:"endDate"`
	Description      string    `json:"description"`
	OverviewLocation string    `json:"overviewLocation"`
	Price            float64   `gorm:"check:price > 0" json:"price"`
	RefundDueDate    string    `json:"refundDueDate"`
	MaxMemberCount   uint      `json:"maxMemberCount"`
	MemberCount      uint      `json:"memberCount"`
	Status           string    `json:"status"`
	AgencyUsername   string    `json:"agencyUsername"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime" json:"createdTimestamp"`
}

func GetTourById(tourId int) (Tour, error) {
	var tour Tour
	result := db.First(&tour, tourId)
	return tour, result.Error
}

func FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo string, offset, limit int) ([]Tour, error) {
	var tours []Tour
	result := db.Model(&Tour{}).Select("tour_id, name, start_date, end_date, overview_location, member_count, max_member_count, price").Where("name LIKE ? AND start_date >= ? AND end_date <= ? AND overview_location LIKE ? AND member_count >= ? AND member_count <= ? AND price >= ? AND price <= ?", name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo).Limit(limit).Offset(offset).Find(&tours)
	return tours, result.Error
}
