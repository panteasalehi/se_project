package controller

import model "MelkOnline/internal/core"

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserID  int    `json:"userid"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ADregisterRequest struct {
	Title              string  `json:"title"`
	Category           string  `json:"category"`
	Price              int     `json:"price"`
	Area               float32 `json:"area"`
	NumberOfRooms      int     `json:"numberOfRooms"`
	YearOfConstruction int     `json:"yearOfConstruction"`
	Floor              string  `json:"floor"`
	Description        string  `json:"description"`
	Elevator           bool    `json:"elevator"`
	Store              bool    `json:"store"`
	Parking            bool    `json:"parking"`
	OwnerID            int     `json:"ownerid"`
	Lt                 float32 `json:"lt"`
	Long               float32 `json:"long"`
}

type ADregisterResponse struct {
	Message string `json:"message"`
	ADID    int    `json:"adid"`
}

type ChatSendMessageRequest struct {
	Content    string `json:"content"`
	SenderID   int    `json:"senderid"`
	ReceiverID int    `json:"receiverid"`
	ChatID     int    `json:"chatid"`
}

type ChatSendMessageResponse struct {
	Message   string `json:"message"`
	MessageID int    `json:"messageid"`
}

type ChatGetMessagesRequest struct {
	ChatID int `json:"chatid"`
}

type ChatGetMessagesResponse struct {
	Message  string          `json:"message"`
	Messages []model.Message `json:"messages"`
}

type MainpageResponse struct {
	Message string     `json:"message"`
	Ads     []model.AD `json:"ads"`
}
