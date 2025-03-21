package city

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/cacheDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restclient"
)

type ViaCepAPI struct {
	client *restclient.RestClient
}

func (a *ViaCepAPI) GetCityByZipCode(ctx context.Context, zipCode string) (*ViaCepCityResponse, error) {
	response := restclient.Request[ViaCepCityResponse, any]{
		Ctx:        ctx,
		Client:     a.client,
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("/ws/%s/json", zipCode),
		Cache:      cacheDB.NewCache[ViaCepCityResponse](fmt.Sprintf("postal-code-%s", zipCode), 5*time.Minute),
	}.Call()

	if response.ErrorBody() != nil {
		return nil, fmt.Errorf("error: %+v", response.ErrorBody())
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return response.SuccessBody(), nil
}
