package models

type Admin struct {
	AdminId   uint   `gorm:"primaryKey" json:"admin_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNo   string `json:"phoneNo"`
}
