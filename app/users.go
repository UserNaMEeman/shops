package app

import "time"

type User struct {
	// Id       int    `json:"-" db:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserOrders struct {
	Number string    `json:"number"`
	Status string    `json:"status"`
	Data   time.Time `json:"uploaded_at"`
}
