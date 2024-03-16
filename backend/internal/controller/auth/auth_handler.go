package auth

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/auth"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	as auth.AuthContract
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		as: auth.NewAuthService(),
	}
}

func (lh *LoginHandler) Login(c echo.Context) error {
	lreq, lres := &controller.LoginRequest{}, &controller.LoginResponse{}
	if err := c.Bind(lreq); err != nil {
		lres.Message = "Invalid request"
		return c.JSON(400, lres)
	}
	if err := c.Validate(lreq); err != nil {
		lres.Message = "Invalid request"
		return c.JSON(400, lres)
	}
	session, err := lh.as.Login(lreq.Email, lreq.Password)
	if err != nil {
		lres.Message = "Invalid credentials"
		return c.JSON(401, lres)
	}
	lres.Message = "Login successful"
	lres.Token = session.Token
	return c.JSON(200, lres)
}
