package models

type Tourist struct {
	Username       string `gorm:"primary_key" json:"username"`
	CitizenId      string `gorm:"unique;not null" json:"citizenId"`
	FirstName      string `gorm:"not null" json:"firstName"`
	LastName       string `gorm:"not null" json:"lastName"`
	Address        string `gorm:"not null" json:"address"`
	BirthDate      string `gorm:"not null" json:"birthDate"`
	Gender         string `gorm:"not null" json:"Gender"`
	DefaultPayment string `gorm:"not null" json:"defaultPayment"`
}

type TouristWithUser struct {
	Username       string `gorm:"primary_key" json:"username"`
	Phone          string `gorm:"not null" json:"phone"`
	Email          string `gorm:"not null" json:"email"`
	Image          string `gorm:"not null" json:"image"`
	Role           string `gorm:"not null" json:"role"`
	CitizenId      string `gorm:"unique;not null" json:"citizenId"`
	FirstName      string `gorm:"not null" json:"firstName"`
	LastName       string `gorm:"not null" json:"lastName"`
	Address        string `gorm:"not null" json:"address"`
	BirthDate      string `gorm:"not null" json:"birthDate"`
	Gender         string `gorm:"not null" json:"Gender"`
	DefaultPayment string `gorm:"not null" json:"defaultPayment"`
}

func ToTouristWithUser(tourist Tourist, user User) TouristWithUser {
	return TouristWithUser{
		Username:       user.Username,
		Phone:          user.Phone,
		Email:          user.Email,
		Image:          user.Image,
		Role:           user.Role,
		CitizenId:      tourist.CitizenId,
		FirstName:      tourist.FirstName,
		LastName:       tourist.LastName,
		Address:        tourist.Address,
		BirthDate:      tourist.BirthDate,
		Gender:         tourist.Gender,
		DefaultPayment: tourist.DefaultPayment,
	}
}
