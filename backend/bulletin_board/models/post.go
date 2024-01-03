package models

import (
	"time"
)

type Post struct {
	ID        int64
	Title     string
	Content   string
	Author    User
	Likes     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
