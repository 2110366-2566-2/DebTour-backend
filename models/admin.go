package models

type Admin struct {
	//create username as a foreign key
	Username  string `gorm:"primary_key" json:"username"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	PhoneNo   string `gorm:"not null" json:"phoneNo"`
}
