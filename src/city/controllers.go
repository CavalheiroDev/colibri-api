package city

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CityController struct {
	useCase ICityUseCases
}

func (c *CityController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "city",
			Method:   http.MethodGet,
			Function: c.List,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "city",
			Method:   http.MethodPost,
			Function: c.Create,
			Prefix:   restserver.PublicApi,
		},
	}
}

func (c *CityController) List(wctx restserver.WebContext) {
	var queryParams CityListQueryParams
	if err := wctx.DecodeQueryParams(&queryParams); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	response, err := c.useCase.List(wctx.Context(), &queryParams)
	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, response)
}

func (c *CityController) Create(wctx restserver.WebContext) {
	var request CityCreateRequest
	if err := wctx.DecodeBody(&request); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	response, err := c.useCase.Create(wctx.Context(), &request)
	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, response)
}
