package chat

import (
	model "MelkOnline/internal/controller"
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/chat"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	cs chat.ChatContract
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		cs: chat.NewChatService(),
	}
}

func (ch *ChatHandler) SendMessage(c echo.Context) error {
	csreq, csres := &model.ChatSendMessageRequest{}, &model.ChatSendMessageResponse{}
	if err := c.Bind(csreq); err != nil {
		csres.Message = "Invalid request"
		return c.JSON(400, csres)
	}
	if err := c.Validate(csreq); err != nil {
		csres.Message = "Invalid request"
		return c.JSON(400, csres)
	}
	message := core.Message{
		Content:    csreq.Content,
		SenderID:   csreq.SenderID,
		ReceiverID: csreq.ReceiverID,
		ChatID:     csreq.ChatID,
	}
	ID, err := ch.cs.SendMessage(message)
	if err != nil {
		csres.Message = "Failed to send message"
		return c.JSON(http.StatusBadRequest, csres)
	}
	csres.Message = "Message sent successfully"
	csres.MessageID = ID
	return c.JSON(200, csres)
}

func (ch *ChatHandler) GetMessagesByChatID(c echo.Context) error {
	cgreq, cgres := &model.ChatGetMessagesRequest{}, &model.ChatGetMessagesResponse{}
	if err := c.Bind(cgreq); err != nil {
		return c.JSON(400, cgres)
	}
	if err := c.Validate(cgreq); err != nil {
		return c.JSON(400, cgres)
	}
	messages, err := ch.cs.GetMessagesByChatID(cgreq.ChatID)
	if err != nil {
		return c.JSON(500, cgres)
	}
	cgres.Message = "Messages retrieved successfully"
	cgres.Messages = messages
	return c.JSON(200, messages)
}
