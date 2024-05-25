package getpost

import (
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
	token, err := c.Cookie("token")
	if err != nil {
		return c.JSON(400, "Token not found")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	ad, err := gph.svc.GetPost(idInt, token.Value)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, ad)
}
