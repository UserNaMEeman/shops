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
	Accrual string    `json:"accrual,omitempty"`
	Date    time.Time `json:"uploaded_at"`
}

type Accruals struct {
	Order   string `json:"order,omitempty"`
	Status  string `json:"status,omitempty"`
	Accrual string `json:"accrual,omitempty"`
}
