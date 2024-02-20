package models

import (
	"time"
)

type User struct {
	Username    string    `gorm:"primary_key" json:"username"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Image       string    `json:"image"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
}
