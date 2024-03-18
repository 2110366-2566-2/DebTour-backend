package models

type Admin struct {
	//create username as a foreign key
	AdminId   uint   `gorm:"primaryKey" json:"admin_id"`
	Username  string `gorm:"not null" json:"username"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	PhoneNo   string `gorm:"not null" json:"phoneNo"`
}
