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

type TouristWithUserAndToken struct {
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
	Token          string `json:"token"`
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

func ToTouristWithUserAndToken(touristWithUser TouristWithUser, token string) TouristWithUserAndToken {
	return TouristWithUserAndToken{
		Username:  touristWithUser.Username,
		Phone:     touristWithUser.Phone,
		Email:     touristWithUser.Email,
		Image:     touristWithUser.Image,
		Role:      touristWithUser.Role,
		CitizenId: touristWithUser.CitizenId,
		FirstName: touristWithUser.FirstName,
		LastName:  touristWithUser.LastName,
		Address:   touristWithUser.Address,
		BirthDate: touristWithUser.BirthDate,
		Token:     token,
	}
}

func ToTourist(touristWithUser TouristWithUser) Tourist {
	return Tourist{
		Username:  touristWithUser.Username,
		CitizenId: touristWithUser.CitizenId,
		FirstName: touristWithUser.FirstName,
		LastName:  touristWithUser.LastName,
		Address:   touristWithUser.Address,
		BirthDate: touristWithUser.BirthDate,
	}
}

func ToUserFromTouristWithUser(touristWithUser TouristWithUser) User {
	return User{
		Username: touristWithUser.Username,
		Phone:    touristWithUser.Phone,
		Email:    touristWithUser.Email,
		Image:    touristWithUser.Image,
		Role:     touristWithUser.Role,
	}
}
