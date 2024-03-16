package main

import (
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/signup"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/signup", signup.SignupHandler)
	e.POST("/login", auth.LoginHandler)
	e.Logger.Fatal(e.Start(":443"))
}
