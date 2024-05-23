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
}

type ADregisterRequest struct {
	Title              string `json:"title"`
	Category           string `json:"category"`
	Price              string `json:"price"`
	Area               string `json:"area"`
	NumberOfRooms      string `json:"numberOfRooms"`
	YearOfConstruction string `json:"yearOfConstruction"`
	Floor              string `json:"floor"`
	Description        string `json:"description"`
	Elevator           string `json:"elevator"`
	Store              string `json:"store"`
	Parking            string `json:"parking"`
	OwnerID            string `json:"ownerid"`
	Lt                 string `json:"lt"`
	Long               string `json:"long"`
	AvatarURL          string `json:"avatarURL"`
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
