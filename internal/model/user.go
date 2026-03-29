package model

import "time"

type User struct {
	ID         uint
	FullName   string    `json:"fullname"`
	Username   string    `json:"username"`
	UserRole   string    `json:"user_role"`
	UserStatus string    `json:"user_status"`
	Password   string    `json:"password"`
	LastLogin  time.Time `json:"last_login"`
	CreatedAT  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
