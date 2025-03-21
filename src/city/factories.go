package city

type CityFactory struct{}

func NewCityRepository() *CityDBRepository {
	return &CityDBRepository{}
}

func NewCityUseCases() *CityUseCases {
	return &CityUseCases{
		cityRepository: NewCityRepository(),
	}
}

func NewCityController() *CityController {
	return &CityController{
		useCase: NewCityUseCases(),
	}
}
