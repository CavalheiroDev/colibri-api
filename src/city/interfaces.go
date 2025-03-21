package city

import (
	"context"
)

type ICityUseCases interface {
	List(ctx context.Context, params *CityListQueryParams) (CityPageResponse, error)
	Create(ctx context.Context, dto *CityCreateRequest) (*CityResponse, error)
}

type ICityRepository interface {
	List(ctx context.Context, params *CityListQueryParams) (CityPageResponse, error)
	Create(ctx context.Context, dto *CityCreateRequest) (*CityResponse, error)
	ExistsByUniqueKey(ctx context.Context, name string, uf string) (*bool, error)
}
