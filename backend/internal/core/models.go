package core

import "time"

type User struct {
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Salt     string  `json:"salt"`
	Name     string  `json:"name"`
	Type     string  `json:"type" gorm:"default:'buyer'"` // buyer, owner
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
	Area               float64 `json:"area"`
	NumberOfRooms      int     `json:"numberOfRooms"`
	YearOfConstruction int     `json:"YearOfConstruction"`
	Floor              string  `json:"floor"`
	Description        string  `json:"description"`
	Elevator           bool    `json:"elevator"`
	Store              bool    `json:"store"`
	Parking            bool    `json:"parking"`
	UserID             int     `json:"userid"`
	Lt                 float64 `json:"lt"`
	Long               float64 `json:"long"`
}

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	SenderID   int       `json:"senderid"`
	ReceiverID int       `json:"receiverid"`
	ChatID     int       `json:"chatid"`
	DateTime   time.Time `json:"datetime"`
}

type Chat struct {
	ID      int `json:"id"`
	User1ID int `json:"user1id"`
	User2ID int `json:"user2id"`
	AdID    int `json:"adid"`
}

type UserNote struct {
	ID     int       `json:"id"`
	UserID int       `json:"userid"`
	ADID   int       `json:"adid"`
	Note   string    `json:"note"`
	Time   time.Time `json:"time"`
}

type Payment struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userid"`
	Amount   int       `json:"amount"`
	Status   string    `json:"status"`
	DateTime time.Time `json:"datetime"`
}
