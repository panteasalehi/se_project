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

type AD struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Category           string  `json:"category"`
	Price              int     `json:"price"`
	Area               float32 `json:"area"`
	NumberOfRooms      int     `json:"numberOfRooms"`
	YearOfConstruction int     `json:"YearOfConstruction"`
	Floor              string  `json:"floor"`
	Description        string  `json:"description"`
	Elevator           bool    `json:"elevator"`
	Store              bool    `json:"store"`
	Parking            bool    `json:"parking"`
	OwnerID            int     `json:"ownerid"`
	// Location
}
