package core

import "time"

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
	UserID             int     `json:"userid"`
	Lt                 float32 `json:"lt"`
	Long               float32 `json:"long"`
	AvatarURL          string  `json:"avatarurl"`
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
	Note   string    `json:"note"`
	Time   time.Time `json:"time"`
}
