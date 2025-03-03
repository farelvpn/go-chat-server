package models

import "time"

type Chat struct {
	ID        int
	UserID    int
	Message   string
	IsReplied bool
	CreatedAt time.Time
}