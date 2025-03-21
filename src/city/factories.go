package city

import (
	"os"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restclient"
)

type CityFactory struct{}

func NewCityRepository() *CityDBRepository {
	return &CityDBRepository{}
}

func NewViaCepAPI() *ViaCepAPI {
	return &ViaCepAPI{
		client: restclient.NewRestClient(&restclient.RestClientConfig{
			Name:    "via-cep-rest-client",
			BaseURL: os.Getenv("VIA_CEP_BASE_URL"),
			Timeout: 10,
		}),
	}
}

func NewCityUseCases() *CityUseCases {
	return &CityUseCases{
		cityRepository: NewCityRepository(),
		cityAPI:        NewViaCepAPI(),
	}
}

func NewCityController() *CityController {
	return &CityController{
		useCase: NewCityUseCases(),
	}
}
