package app

type User struct {
	// Id       int    `json:"-" db:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
