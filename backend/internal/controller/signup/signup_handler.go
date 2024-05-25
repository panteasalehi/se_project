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

// Signup	 	 user signup
//
//	@Summary		user signup
//	@Description	get user's email and password and name and signup
//	@Tags			signup
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"user's Email"
//	@Param			Password	body		string	true	"user's password"
//	@Param			Name		body		string	true	"user's password"
//	@Success		200			{string}	string
//	@Failure		400			{string}	string
//	@Failure		500			{string}	string
//	@Router			/signup/ [post]
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
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, sres)
}
