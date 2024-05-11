package main

import (
	"MelkOnline/internal/controller/ADregister"
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/signup"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/signup", signup.NewSignupHandler().Signup)
	e.POST("/login", auth.NewLoginHandler().Login)
	e.POST("/ADregister", ADregister.NewADregisterHandler().ADregister)
	e.Logger.Fatal(e.Start(":8086"))
}
