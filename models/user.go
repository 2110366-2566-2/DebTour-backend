package models

import (
	"time"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	CreatedTime time.Time
}
