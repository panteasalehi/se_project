package mainpage

import (
	model "MelkOnline/internal/controller"
	Mainpage "MelkOnline/internal/core/mainpage"

	"github.com/labstack/echo/v4"
)

type MainpageHandler struct {
	cs Mainpage.MainpageContract
}

func NewMainpageHandler() *MainpageHandler {
	return &MainpageHandler{
		cs: Mainpage.NewMainpageService(),
	}
}

func (m *MainpageHandler) GetAds(c echo.Context) error {
	mres := &model.MainpageResponse{}
	session, err := c.Cookie("session")
	if err != nil {
		mres.Message = "Session not found"
		return c.JSON(400, mres)
	}
	Ads, err := m.cs.Mainpage(session.Value)
	if err != nil {
		mres.Message = err.Error()
		return c.JSON(500, mres)
	}
	mres.Message = "Ads retrieved successfully"
	mres.Ads = Ads
	return c.JSON(200, Ads)
}
