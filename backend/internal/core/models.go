package core

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type Session struct {
	Token string `json:"token"`
}
