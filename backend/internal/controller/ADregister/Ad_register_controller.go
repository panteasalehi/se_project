package ADregister

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/ADregister"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ADregisterHandler struct {
	ss ADregister.ADregisterContract
}

func NewADregisterHandler() *ADregisterHandler {
	return &ADregisterHandler{
		ss: ADregister.NewADregisterService(),
	}
}

func (adh *ADregisterHandler) ADregister(c echo.Context) error {
	adreq := &controller.ADregisterRequest{}
	adres := &controller.ADregisterResponse{}
	adreqPriceInt, err1 := strconv.Atoi(c.Request().FormValue("price"))
	adreqAreaFloat, err2 := strconv.ParseFloat(c.Request().FormValue("area"), 64)
	adreqNumberOfRoomsInt, err3 := strconv.Atoi(c.Request().FormValue("numberOfRooms"))
	adreqYearOfConstructionInt, err4 := strconv.Atoi(c.Request().FormValue("yearOfConstruction"))
	adreqLtFloat, err5 := strconv.ParseFloat(c.Request().FormValue("lt"), 64)
	adreqLongFloat, err6 := strconv.ParseFloat(c.Request().FormValue("long"), 64)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		adres.Message = err1.Error()
		return c.JSON(http.StatusInternalServerError, adres)
	}
	var err7, err8, err9 error
	adreq.Title = c.Request().FormValue("title")
	adreq.Category = c.Request().FormValue("category")
	adreq.Floor = c.Request().FormValue("floor")
	adreq.Description = c.Request().FormValue("description")
	adreq.Elevator, err7 = strconv.ParseBool(c.Request().FormValue("elevator"))
	adreq.Store, err8 = strconv.ParseBool(c.Request().FormValue("store"))
	adreq.Parking, err9 = strconv.ParseBool(c.Request().FormValue("parking"))
	adreq.AvatarURL = c.Request().FormValue("avatarURL")
	if err7 != nil || err8 != nil || err9 != nil {
		adres.Message = err7.Error()
		return c.JSON(http.StatusInternalServerError, adres)
	}
	cookie, _ := c.Cookie("session")
	token := cookie.Value
	ID, err := adh.ss.ADregister(token, adreq.Title, adreq.Category, adreqPriceInt, adreqAreaFloat, adreqNumberOfRoomsInt, adreqYearOfConstructionInt, adreq.Floor, adreq.Description, adreq.Elevator, adreq.Store, adreq.Parking, 0, adreqLtFloat, adreqLongFloat, adreq.AvatarURL)
	if err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, adres)
	}
	adres.Message = "Advertisement registration completed"
	adres.ADID = ID
	return c.JSON(http.StatusOK, adres)
}
