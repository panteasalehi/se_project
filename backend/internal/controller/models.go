package controller

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignupResponse struct {
	Message string `json:"message"`
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
	// Location
}

type ADregisterResponse struct {
	Message string `json:"message"`
	//?
}

type ChatSendMessageRequest struct {
	Content    string `json:"content"`
	SenderID   int    `json:"senderid"`
	ReceiverID int    `json:"receiverid"`
	ChatID     int    `json:"chatid"`
}

type ChatSendMessageResponse struct {
	Message string `json:"message"`
}

type ChatGetMessagesRequest struct {
	ChatID int `json:"chatid"`
}
