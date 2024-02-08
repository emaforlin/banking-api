package models

// add data validations

type AddUserData struct {
	FullName string `json:"fullname"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ShowUserData struct {
	FullName string `json:"fullName"`
	Funds    int64  `json:"funds"`
}
