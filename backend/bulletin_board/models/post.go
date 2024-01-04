package models

import (
	"time"
)

type Post struct {
	ID        int64
	Title     string
	Content   string
	Author    *User `orm:"rel(fk)"` // 외래 키로 User 모델을 참조
	Likes     int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
