package chat

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/chat"
	"MelkOnline/internal/core/getpost"
	"MelkOnline/internal/infrastructure/auth"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebsocketData struct {
	Connection *websocket.Conn
	SenderID   int
	ReceiverID int
	ChatID     int
}

var upgrader = websocket.Upgrader{}
var connections = make(map[int]map[int]*WebsocketData)
var connMutex = &sync.Mutex{}

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

// SendMessage godoc
//
//	@Summary		Send message
//	@Description	Send message to the user
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			ad_id	path		int		true	"Ad ID"
//	@Success		200		{string}	string	"message sent"
//	@Failure		400		{string}	string	"error in sending message"
//	@Router			/api/v1/ads/{ad_id}/chats [post]
func (ch *ChatHandler) SendMessage(c echo.Context) error {
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	userID, err := ch.sr.ValidateSession(token.Value)
	if err != nil {
		return c.JSON(400, "error in validating session")
	}
	var message core.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(400, "error in binding message")
	}
	adID := c.Param("ad_id")
	adIDint, _ := strconv.Atoi(adID)
	connMutex.Lock()
	for _, conn := range connections[adIDint] {
		if conn.SenderID == userID {
			message.ReceiverID = conn.ReceiverID
		} else if conn.ReceiverID == userID {
			message.ReceiverID = conn.SenderID
		}
	}
	conn := connections[adIDint][message.SenderID]
	connMutex.Unlock()
	conn.Connection.WriteJSON(message)
	message.SenderID = userID
	message.ChatID, _ = ch.cs.ChatExists(adIDint, userID)
	message.ChatID = conn.ChatID
	_, err = ch.cs.SendMessage(message)
	if err != nil {
		return c.JSON(400, "error in sending message")
	}
	return c.JSON(200, "message sent")
}

// GetMessage godoc
//
//	@Summary		Get message
//	@Description	Get message from the user
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			ad_id	path		int		true	"Ad ID"
//	@Success		200		{string}	string	"message received"
//	@Failure		400		{string}	string	"error in reading message"
//	@Router			/api/v1/ads/{ad_id}/chats [get]
func (ch *ChatHandler) GetMessage(c echo.Context) error {
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	userID, err := ch.sr.ValidateSession(token.Value)
	if err != nil {
		return c.JSON(400, "error in validating session")
	}
	adID := c.Param("ad_id")
	adIDint, _ := strconv.Atoi(adID)
	ad, _ := ch.gc.GetPost(adIDint, token.Value)
	chatid, _ := ch.cs.ChatExists(adIDint, userID)
	if chatid == 0 {
		chatid, err = ch.cs.CreateChat(ad.UserID, userID, adIDint)
		if err != nil {
			return c.JSON(400, "error in creating chat")
		}
	}
	connMutex.Lock()
	if _, ok := connections[adIDint][userID]; !ok {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return c.JSON(400, "error in websocket connection")
		}
		connections[adIDint][userID] = &WebsocketData{
			Connection: conn,
			ReceiverID: ad.UserID,
			SenderID:   userID,
			ChatID:     chatid,
		}
	}
	conn := connections[adIDint][userID]
	connMutex.Unlock()
	for {
		var message core.Message
		err := conn.Connection.ReadJSON(&message)
		if err != nil {
			return c.JSON(400, "error in reading message")
		}
		c.String(200, message.Content)
	}
}
