package models

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Role      string
	OTP       string
	CreatedAt time.Time
}