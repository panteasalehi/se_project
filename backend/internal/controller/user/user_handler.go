package user

import (
	"MelkOnline/internal/core/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	us user.UserServiceContract
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		us: user.NewUserService(),
	}
}

func (uh *UserHandler) GetUserBySession(c echo.Context) error {
	token, err := c.Request().Cookie("token")
	if err != nil {
		return c.JSON(400, err.Error())
	}
	user, err := uh.us.GetUserBySession(token.Value)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, user)
}

func (uh *UserHandler) GetUserAds(c echo.Context) error {
	token, err := c.Request().Cookie("token")
	if err != nil {
		return c.JSON(400, err.Error())
	}
	ads, err := uh.us.GetUserAds(token.Value)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, ads)
}
