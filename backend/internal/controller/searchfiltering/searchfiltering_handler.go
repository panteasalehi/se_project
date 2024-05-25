package searchfiltering

import (
	model "MelkOnline/internal/controller"
	"MelkOnline/internal/core"
	searchfiltering "MelkOnline/internal/core/searchfiltering"

	"net/http"

	"github.com/labstack/echo/v4"
)

type SearchFilteringHandler struct {
	ss searchfiltering.SearchFilteringContract
}

func NewSearchFilteringHandler() *SearchFilteringHandler {
	return &SearchFilteringHandler{
		ss: searchfiltering.NewSearchFilteringService(),
	}
}

func unique(arr []core.AD) []core.AD {
	uniqueElements := make(map[core.AD]bool)
	result := make([]core.AD, 0, len(arr))
	for _, num := range arr {
		if !uniqueElements[num] {
			uniqueElements[num] = true
			result = append(result, num)
		}
	}
	return result
}

// Searchfiltering	Searchfiltering
//
//	@Summary		Searchfiltering
//	@Description	filter ads and get the desired ads from database
//	@Tags			searchfiltering
//	@Accept			json
//	@Produce		json
//	@Param			Category			body		string		true	"Category to filter"
//	@Param			MaxPrice			body		int			true	"Max Price to filter"
//	@Param			MinPrice			body		int			true	"Min Price to  filter"
//	@Param			MaxArea				body		int			true	"Max Area to filter"
//	@Param			MinArea				body		int			true	"Min Area to filter"
//	@Param			NumberOfRooms		body		int			true	"Number of Rooms to filter"
//	@Param			MaxAge				body		int			true	"Max age to filter"
//	@Param			MaxFloor			body		int			true	"Max Floor to filter"
//	@Param			MinFloor			body		int			true	"Min Floor to filter"
//	@Param			Elevator			body		bool		true	"availability of Elevator"
//	@Param			Store				body		bool		true	"availability of Store"
//	@Param			Parking				body		bool		true	"availability of Parking"
//	@Success		200			{string}	string
//	@Failure		400			{string}	string
//	@Failure		500			{string}	string
//	@Router			/searchfiltering/{Category,MaxPrice,MinPrice,MaxArea,MinArea,NumberOfRooms,MaxAge,MaxFloor,MinFloor,Elevator,Store,Parking} [get]
func (sfh *SearchFilteringHandler) Searchfiltering(c echo.Context) error {
	sfreq := &model.SearchFilteringRequest{}
	sfres := &model.SearchFilteringResponse{}
	if err := c.Bind(sfreq); err != nil {
		sfres.Message = err.Error()
		return c.JSON(http.StatusBadRequest, sfres)
	}

	Ads, err := sfh.ss.FindCategoryFilteredADs(sfreq.Category)
	Ads1, err := sfh.ss.FindPriceFilteredADs(sfreq.MinPrice, sfreq.MaxPrice)
	Ads2, err := sfh.ss.FindAreaFilteredADs(sfreq.MinArea, sfreq.MaxArea)
	Ads3, err := sfh.ss.FindRoomFilteredADs(sfreq.NumberOfRooms)
	Ads4, err := sfh.ss.FindPSEFilteredADs(sfreq.Parking, sfreq.Store, sfreq.Elevator)
	Ads5, err := sfh.ss.FindFloorFilteredADs(sfreq.MinFloor, sfreq.MaxFloor)
	Ads6, err := sfh.ss.FindYearFilteredADs(sfreq.MaxAge)

	result := append(append(append(append(append(append(Ads, Ads1...), Ads2...), Ads3...), Ads4...), Ads5...), Ads6...)

	result = unique(result)

	if err != nil {
		sfres.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, sfres)
	}
	sfres.Message = "search filters applied"
	sfres.Ads = result
	return c.JSON(http.StatusOK, sfres)
}
