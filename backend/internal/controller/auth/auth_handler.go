package auth

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/auth"
	"net/http"

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

// Login	 	 user Login
//
//	@Summary		user Login
//	@Description	get user's email and password and Login
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"user's Email"
//	@Param			Password	body		string	true	"user's password"
//	@Success		200			{string}	string
//	@Failure		400			{string}	string
//	@Failure		401			{string}	string
//	@Failure		500			{string}	string
//	@Router			/login/ [post]
func (lh *LoginHandler) Login(c echo.Context) error {
	lreq, lres := &controller.LoginRequest{}, &controller.LoginResponse{}
	if err := c.Bind(lreq); err != nil {
		lres.Message = "Invalid request"
		return c.JSON(400, lres)
	}
	session, err := lh.as.Login(lreq.Email, lreq.Password)
	if err != nil {
		lres.Message = err.Error()
		return c.JSON(401, lres)
	}
	lres.Message = "Login successful"
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = session.Token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)
	return c.JSON(200, lres)
}
