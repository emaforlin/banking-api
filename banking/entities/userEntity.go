package entities

import "time"

type (
	InsertUserDto struct {
		Id        uint64    `gorm:"primaryKey,autoincrement" json:"id"`
		FullName  string    `json:"fullname"`
		Username  string    `gorm:"not null; unique" json:"username"`
		Password  string    `gorm:"not null" json:"password"`
		Funds     int64     `json:"funds"`
		CreatedAt time.Time `json:"createdAt"`
	}

	User struct {
		Id        uint64    `json:"id"`
		FullName  string    `json:"fullName"`
		Username  string    `json:"user"`
		Password  string    `json:"password"`
		Funds     int64     `json:"funds"`
		CreatedAt time.Time `json:"createdAt"`
	}
)
