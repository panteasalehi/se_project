package chat

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/chat"
	"MelkOnline/internal/core/getpost"
	"MelkOnline/internal/infrastructure/auth"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}
var connections = make(map[int]map[int]*websocket.Conn)

type ChatHandler struct {
	cs chat.ChatContract
	gc getpost.GetPostContract
	sr *auth.SessionRepository
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		cs: chat.NewChatService(),
		gc: getpost.NewGetPostService(),
		sr: auth.NewSessionRepository(),
	}
}

// SendMessage	 send messages in chat
//
//	@Summary		send messages in chat
//	@Description	gets content of message and reciever and sender and chat's ID and sending the message to reciever
//	@Tags			Chat
//	@Accept			json
//	@Produce		json
//	@Param			chatid	body		int	true	"Chat's ID"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string
//	@Failure		400		{string}	string
//	@Failure		400		{string}	string
//	@Router			/Chat/send/ [post]
func (ch *ChatHandler) SendMessage(c echo.Context) error {
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	message := core.Message{}
	if err := c.Bind(&message); err != nil {
		return c.JSON(400, "error in binding message")
	}
	_, err = ch.sr.ValidateSession(token.Value)
	if err != nil {
		return c.JSON(400, "error in validating session")
	}
	conn := connections[message.ReceiverID][message.SenderID]
	err = conn.WriteJSON(message)
	if err != nil {
		return c.JSON(400, "error in sending message")
	}
	return c.JSON(200, "message sent")
}

// GetMessagesByChatID	 get messages by chat's ID
//
//	@Summary		get messages by chat's ID
//	@Description	gets chat's ID and returns all messages in that chat
//	@Tags			Chat
//	@Accept			json
//	@Produce		json
//	@Param			chatid	path		int	true	"Chat's ID"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string
//	@Failure		400		{string}	string
//	@Failure		400		{string}	string
//	@Router			/Chat/{chatid} [get]

func (ch *ChatHandler) GetMessage(c echo.Context) error {
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	adID := c.Param("ad_id")
	userID := c.Param("user_id")
	adIDint, _ := strconv.Atoi(adID)
	userIDint, _ := strconv.Atoi(userID)
	ad, err := ch.gc.GetPost(adIDint, token.Value)
	if err != nil {
		return c.JSON(400, "ad not found")
	}
	if exists, _ := ch.cs.ChatExists(adIDint, userIDint); !exists {
		ch.cs.CreateChat(ad.UserID, userIDint, ad.ID)
	}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), http.Header{
		"Access-Control-Allow-Origin": []string{"*"},
	})
	connections[ad.UserID][userIDint] = conn
	if err != nil {
		return c.JSON(400, "error in websocket connection")
	}
	for {
		message := core.Message{}
		err := conn.ReadJSON(&message)
		if err != nil {
			c.JSON(400, "error in reading message")
		}
		message.SenderID = userIDint
		message.ChatID = adIDint
		_, err = ch.cs.SendMessage(message)
		c.JSON(200, message)
		if err != nil {
			c.JSON(400, "error in sending message")
		}
	}
}
