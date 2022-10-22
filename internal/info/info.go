package info

type user struct {
	Login    string
	Password string
}

func NewUser() user {
	user := user{}
	return user
}
