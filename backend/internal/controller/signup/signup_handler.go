package signup

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/signup"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type SignupHandler struct {
	ss signup.SignupContract
}

func NewSignupHandler() *SignupHandler {
	return &SignupHandler{
		ss: signup.NewSignupService(),
	}
}

func (sh *SignupHandler) Signup(c echo.Context) error {
	sreq := &controller.SignupRequest{}
	sres := &controller.SignupResponse{}
	if err := c.Bind(sreq); err != nil {
		sres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, sres)
	}
	ID, token, err := sh.ss.Signup(sreq.Email, sreq.Password, sreq.Name)
	if err != nil {
		sres.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, sres)
	}
	sres.Message = "User created successfully"
	sres.UserID = ID
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, sres)
}
