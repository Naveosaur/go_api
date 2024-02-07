package entity

import "time"

type Posts struct {
	ID         int
	UserID     int
	Post       string
	PictureUrl *string // Agar bisa nil
	CreatedAt  time.Time
}