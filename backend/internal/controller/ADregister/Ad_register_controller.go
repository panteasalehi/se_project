package ADregister

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/ADregister"
	"MelkOnline/internal/core/user"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type ADregisterHandler struct {
	ss ADregister.ADregisterContract
	ur user.UserServiceContract
}

func NewADregisterHandler() *ADregisterHandler {
	return &ADregisterHandler{
		ss: ADregister.NewADregisterService(),
		ur: user.NewUserService(),
	}
}

// ADregister	 registration of ADs
//
//	@Summary		registration of ADs
//	@Description	registration of ADs and giving each an ID
//	@Tags			ADregister
//	@Accept			json
//	@Produce		json
//	@Param			ad						body		controller.ADregisterRequest	true	"Advertisement data"
//	@Success		200						{string}	string
//	@Failure		400						{string}	string
//	@Failure		500						{string}	string
//	@Router			/api/v1/ads/register	[post]
func (adh *ADregisterHandler) ADregister(c echo.Context) error {
	adreq := &controller.ADregisterRequest{}
	adres := &controller.ADregisterResponse{}
	if err := c.Bind(adreq); err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, adres)
	}
	cookie, err := c.Request().Cookie("session")
	if err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, adres)
	}
	if user, err := adh.ur.GetUserBySession(cookie.Value); err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, adres)
	} else if user.Type != "owner" {
		adres.Message = "You are not allowed to register an advertisement"
		return c.JSON(http.StatusBadRequest, adres)
	}
	token := cookie.Value
	file, err := c.FormFile("image")
	if err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, adres)
	}
	ID, err := adh.ss.ADregister(token, adreq.Title, adreq.Category, adreq.Price, adreq.Area, adreq.NumberOfRooms, adreq.YearOfConstruction, adreq.Floor, adreq.Description, adreq.Elevator, adreq.Store, adreq.Parking, adreq.OwnerID, adreq.Lt, adreq.Long, file)
	if err != nil {
		adres.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, adres)
	}
	adres.Message = "Advertisement registration completed"
	adres.ADID = ID
	return c.JSON(http.StatusOK, adres)
}
