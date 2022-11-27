package app

import "time"

type User struct {
	// Id       int    `json:"-" db:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserOrders struct {
	Order   string    `json:"number"`
	Status  string    `json:"status"`
	Accrual float64   `json:"accrual,omitempty"`
	Date    time.Time `json:"uploaded_at"`
}

type Accruals struct {
	Order   string  `json:"order,omitempty"`
	Status  string  `json:"status,omitempty"`
	Accrual float64 `json:"accrual,omitempty"`
}

type Balance struct {
	Current   float64 `json:"current,omitempty"`
	Withdrawn float64 `json:"withdrawn,omitempty"`
}

type Buy struct {
	Order string    `json:"order"`
	Sum   float64   `json:"sum"`
	Date  time.Time `json:"processed_at"`
}
