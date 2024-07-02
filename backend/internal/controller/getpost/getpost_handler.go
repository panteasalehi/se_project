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

// GetPost	 get post by id
//
//	@Summary		get post by id
//	@Description	get post by id
//	@Tags			getpost
//	@Accept			json
//	@Produce		json
//	@Param			ad_id	path		int	true	"post id"
//	@Success		200		{object}	core.AD
//	@Failure		400		{string}	string
//	@Failure		500		{string}	string
//	@Router			/api/v1/ads/{ad_id} [get]
func (gph *GetPostHandler) GetPost(c echo.Context) error {
	id := c.Param("ad_id")
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
