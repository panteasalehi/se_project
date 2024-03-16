package core

type User struct {
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Salt     string  `json:"salt"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Score    float32 `json:"score"`
}

type Session struct {
	Token string `json:"token"`
}
