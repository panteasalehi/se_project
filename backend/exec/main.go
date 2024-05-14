package main

import (
	"MelkOnline/internal/controller/ADregister"
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/chat"
	"MelkOnline/internal/controller/signup"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "public")
	e.File("/", "public/index.html")


	e.POST("/signup", signup.NewSignupHandler().Signup)
	e.POST("/login", auth.NewLoginHandler().Login)
	e.POST("/ADregister", ADregister.NewADregisterHandler().ADregister)
	e.GET("/Chat/page/?chatid=", chat.NewChatHandler().GetMessagesByChatID)
	e.POST("/Chat/send/?chatid=", chat.NewChatHandler().SendMessage)
	
	e.Logger.Fatal(e.Start(":8080"))
}
