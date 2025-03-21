package city

import "github.com/colibri-project-io/colibri-sdk-go/pkg/base/types"

type CityListQueryParams struct {
	Page     uint16 `form:"page" validate:"required"`
	PageSize uint16 `form:"pageSize" validate:"required"`
	Name     string `form:"name"`
	State    string `form:"state"`
}

type CityResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CityPageResponse *types.Page[CityResponse]

type CityCreateRequest struct {
	Name  string `json:"name" validate:"required"`
	State string `json:"state" validate:"required"`
}

type ViaCepCityResponse struct {
	Name  string `json:"localidade"`
	State string `json:"uf"`
}
