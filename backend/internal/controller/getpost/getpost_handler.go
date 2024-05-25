package getpost

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/getpost"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type GetPostHandler struct {
	svc getpost.GetPostContract
}

func NewGetPostHandler() *GetPostHandler {
	return &GetPostHandler{
		svc: getpost.NewGetPostService(),
	}
}

func (gph *GetPostHandler) GetPost(c echo.Context) error {
	id := c.Param("id")
	token, err := c.Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	ad, err := gph.svc.GetPost(idInt, token.Value)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	adArr := make([]core.AD, 1)
	adArr[0] = *ad
	return c.JSON(200, ad)
}
