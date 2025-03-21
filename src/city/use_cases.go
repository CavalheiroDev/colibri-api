package city

import (
	"context"
	"errors"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/logging"
)

type CityUseCases struct {
	cityRepository ICityRepository
	cityAPI        ICityAPI
}

func (u *CityUseCases) List(ctx context.Context, params *CityListQueryParams) (CityPageResponse, error) {
	response, err := u.cityRepository.List(ctx, params)
	if err != nil {
		logging.Error("error in List all cities: %+v", err)
		return nil, errors.New(ErrOnListCity)
	}

	return response, nil
}

func (u *CityUseCases) Create(ctx context.Context, dto *CityCreateRequest) (*CityResponse, error) {
	exists, err := u.cityRepository.ExistsByUniqueKey(ctx, dto.Name, dto.State)
	if err != nil {
		logging.Error("error in ExistsByUniqueKey: %+v", err)
		return nil, errors.New(ErrOnCreateCity)
	}

	if *exists {
		return nil, errors.New(ErrCityAlreadyExists)
	}

	response, err := u.cityRepository.Create(ctx, dto)
	if err != nil {
		logging.Error("error in Create city: %+v", err)
		return nil, errors.New(ErrOnCreateCity)
	}

	return response, nil
}

func (u *CityUseCases) GetCityByZipCode(ctx context.Context, zipCode string) (*ViaCepCityResponse, error) {
	response, err := u.cityAPI.GetCityByZipCode(ctx, zipCode)
	if err != nil {
		logging.Error("error in GetCityByZipCode: %+v", err)
		return nil, errors.New(ErrOnGetCityByZipCode)
	}

	return response, nil
}
