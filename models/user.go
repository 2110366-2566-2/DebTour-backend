package models

import (
	"time"
)

type User struct {
	Username    string    `gorm:"primary_key" json:"username"`
	Password    string    `gorm:"not null" json:"password"`
	Phone       string    `gorm:"not null" json:"phone"`
	Email       string    `gorm:"not null" json:"email"`
	Image       string    `gorm:"not null" json:"image"`
	Role        string    `gorm:"not null" json:"role"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
}
