package models

type Admin struct {
	AdminId   uint   `gorm:"primaryKey" json:"admin_id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	PhoneNo   string `gorm:"not null" json:"phoneNo"`
}
