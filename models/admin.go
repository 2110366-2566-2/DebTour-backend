package models

type Admin struct {
	AdminId   uint   `json:"adminId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNo   string `json:"phoneNo"`
}
