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
	t, _ := c.Cookie("token")
	Ads, err := m.cs.Mainpage(t.Value)
	if err != nil {
		return c.JSON(500, mres)
	}
	mres.Message = "Ads retrieved successfully"
	mres.Ads = Ads
	return c.JSON(200, Ads)
}
