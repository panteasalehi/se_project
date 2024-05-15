package ADregister

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/ADregister"
	"net/http"

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
	if err := c.Bind(adreq); err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, adres)
	}
	ID, err := adh.ss.ADregister(adreq.Title, adreq.Category, adreq.Price, adreq.Area, adreq.NumberOfRooms, adreq.YearOfConstruction,
		adreq.Floor, adreq.Description, adreq.Elevator, adreq.Store, adreq.Parking, adreq.OwnerID)
	if err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, adres)
	}
	adres.Message = "Advertisement registration completed"
	adres.ADID = ID
	return c.JSON(http.StatusOK, adres)
}
