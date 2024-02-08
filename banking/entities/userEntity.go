package entities

import "time"

type (
	InsertUserDto struct {
		Id        uint64    `gorm:"primaryKey,autoincrement" json:"id"`
		FullName  string    `json:"fullName"`
		Funds     int64     `json:"funds"`
		CreatedAt time.Time `json:"createdAt"`
	}

	RetrieveUserDto struct {
		FullName string `json:"fullName"`
		Funds    int64  `json:"funds"`
	}

	User struct {
		Id        uint64    `json:"id"`
		FullName  string    `json:"fullName"`
		Funds     int64     `json:"funds"`
		CreatedAt time.Time `json:"createdAt"`
	}
)
